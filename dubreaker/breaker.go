package dubreaker

import (
	"errors"
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/config"
	"log"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

type Breaker struct {
	Name         string       `json:"name" toml:"name" yaml:"name"`
	LogDir       string       `json:"log_dir" toml:"log_dir" yaml:"log_dir"`
	SlowReqRule  *BreakerRule `json:"slow_req_rule" toml:"slow_req_rule" yaml:"slow_req_rule"`
	ErrRatioRule *BreakerRule `json:"err_ratio_rule" toml:"err_ratio_rule" yaml:"err_ratio_rule"`
	DingTalk     bool         `json:"ding_talk" yaml:"ding_talk"` // 熔断发送钉钉通知，默认关闭
}

type BreakerRule struct {
	Strategy         circuitbreaker.Strategy `json:"strategy" toml:"strategy" yaml:"strategy"`                               // 熔断策略 0：慢请求比率 1：错误请求比率 2：错误量
	RetryTimeoutMs   uint32                  `json:"retry_timeout_ms" toml:"retry_timeout_ms" yaml:"retry_timeout_ms"`       // 即熔断触发后持续的时间（ms),结束后切换到 half-open
	MinRequestAmount uint64                  `json:"min_request_amount" toml:"min_request_amount" yaml:"min_request_amount"` // 静默数量，如果当前统计周期内对资源的访问数量小于静默数量，那么熔断器就处于静默期。
	StatIntervalMs   uint32                  `json:"stat_interval_ms" toml:"stat_interval_ms" yaml:"stat_interval_ms"`       // 统计的时间窗口长度（单位为 ms）
	Threshold        float64                 `json:"threshold" toml:"threshold" yaml:"threshold"`                            // Threshold表示是慢调用比例的阈值(小数表示，比如0.1表示10%)，也就是如果当前资源的慢调用比例如果高于Threshold，那么熔断器就会断开；否则保持闭合状态。 对于错误比例策略，Threshold表示的是错误比例的阈值(小数表示，比如0.1表示10%)。对于错误数策略，Threshold是错误计数的阈值。
	MaxAllowedRtMs   uint64                  `json:"max_allowed_rt_ms" toml:"max_allowed_rt_ms" yaml:"max_allowed_rt_ms"`    // 最大执行时间（ms 只对满请求起作用）
}

// Entry 熔断器保护入口（有返回值）
// operate: 受保护的组件
// resultWithBreaker：熔断发生后期望响应的默认值
// errs：错误熔断策略中期望不受监控的错误
func (b *Breaker) Entry(operate func() (interface{}, error), resultWithBreaker interface{}, errs ...error) (interface{}, error) {
	return wrapEntry(b.Name, operate, resultWithBreaker, errs...)
}

// EntryWithoutResult 熔断器保护入口（无返回值）
// operate: 受保护的组件
// resultWithBreaker：熔断发生后期望响应的默认值
// errs：错误熔断策略中期望不受监控的错误
func (b *Breaker) EntryWithoutResult(operate func() error, errs ...error) error {
	return wrapEntryWithoutResult(b.Name, operate, errs...)
}

// InitBreaker 初始化熔断实例
// logDir：存放熔断相关的日志目录
// busName：受保护的资源名字
// slowReqRule：慢请求熔断策略，可以为 nil
// errRatioRule：错误请求熔断策略，可以为 nil
// dingTalk: 发生熔断是否发送钉钉通知
func InitBreaker(logDir, busName string, slowReqRule, errRatioRule *BreakerRule, dingTalk bool) (*Breaker, error) {
	if slowReqRule == nil && errRatioRule == nil {
		return nil, errors.New("no circuit breaker rules set")
	}
	// 1、对 Sentinel 的运行环境进行相关配置并初始化
	conf := config.NewDefaultConfig()
	if logDir == "" || busName == "" {
		return nil, errors.New("log path or business name can not be empty")
	}
	c := &Breaker{
		Name:         busName,
		LogDir:       logDir,
		SlowReqRule:  slowReqRule,
		ErrRatioRule: errRatioRule,
		DingTalk:     dingTalk,
	}

	conf.Sentinel.Log.Dir = c.LogDir
	conf.Sentinel.App.Name = c.Name

	err := sentinel.InitWithConfig(conf)
	if err != nil {
		return nil, err
	}
	loaRules := make([]*circuitbreaker.Rule, 0)
	if c.ErrRatioRule != nil {
		loaRules = append(loaRules, &circuitbreaker.Rule{
			Resource:         c.Name,
			Strategy:         c.ErrRatioRule.Strategy,
			RetryTimeoutMs:   c.ErrRatioRule.RetryTimeoutMs,
			MinRequestAmount: c.ErrRatioRule.MinRequestAmount,
			StatIntervalMs:   c.ErrRatioRule.StatIntervalMs,
			Threshold:        c.ErrRatioRule.Threshold,
		})
	}
	if c.SlowReqRule != nil {
		loaRules = append(loaRules, &circuitbreaker.Rule{
			Resource:         c.Name,
			Strategy:         c.SlowReqRule.Strategy,
			RetryTimeoutMs:   c.SlowReqRule.RetryTimeoutMs,
			MinRequestAmount: c.SlowReqRule.MinRequestAmount,
			StatIntervalMs:   c.SlowReqRule.StatIntervalMs,
			Threshold:        c.SlowReqRule.Threshold,
			MaxAllowedRtMs:   c.SlowReqRule.MaxAllowedRtMs,
		})
	}
	if len(loaRules) == 0 {
		return nil, errors.New("no circuit breaker rules set")
	}
	// 2、埋点（定义资源），该步骤主要是确定系统中有哪些资源需要防护
	// 资源可以是应用、接口、函数、甚至是一段代码
	_, err = circuitbreaker.LoadRules(loaRules)
	if err != nil {
		return nil, err
	}
	if c.DingTalk {
		// Register a state change listener so that we could observer the state change of the internal circuit breaker.
		circuitbreaker.RegisterStateChangeListeners(&stateChangeTestListener{Name: c.Name})
	}
	return c, nil
}

type stateChangeTestListener struct {
	Name string
}

// OnTransformToClosed 规则：abc_errorRatio   切换前：HalfOpen        切换后：Closed  时间: 2021-06-10 19:22:46
func (s *stateChangeTestListener) OnTransformToClosed(prev circuitbreaker.State, rule circuitbreaker.Rule) {

	msg := fmt.Sprintf("  \n  规则：%s  \n  切换前：%s  \n  切换后：Closed  \n  时间: %s", rule.Resource+"_"+rule.Strategy.String(), prev.String(), time.Now().Format(TimeFormat))
	send("熔断降级", s.Name, msg)
	log.Println(msg)
}

// OnTransformToOpen 规则：abc_errorRatio    切换前：Closed  切换后：Open    触发值: 0.3333333333333333      时间: 2021-06-10 19:21:55
func (s *stateChangeTestListener) OnTransformToOpen(prev circuitbreaker.State, rule circuitbreaker.Rule, snapshot interface{}) {

	msg := fmt.Sprintf("  \n  规则：%s  \n  切换前：%s  \n  切换后：Open  \n  触发值: %+v  \n  时间: %s", rule.Resource+"_"+rule.Strategy.String(), prev.String(), snapshot, time.Now().Format(TimeFormat))
	send("熔断降级", s.Name, msg)
	log.Println(msg)
}

// OnTransformToHalfOpen 规则：abc_errorRatio   切换前：Open    切换后：Half-Open       时间: 2021-06-10 19:21:58
func (s *stateChangeTestListener) OnTransformToHalfOpen(prev circuitbreaker.State, rule circuitbreaker.Rule) {

	msg := fmt.Sprintf("  \n  规则：%s  \n  切换前：%s  \n  切换后：Half-Open  \n  时间: %s", rule.Resource+"_"+rule.Strategy.String(), prev.String(), time.Now().Format(TimeFormat))
	send("熔断降级", s.Name, msg)
	log.Println(msg)
}
