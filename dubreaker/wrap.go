package dubreaker

import (
	"github.com/alibaba/sentinel-golang/api"
)

// wrapEntry 熔断装饰函数
// resource 代表资源名字
// operate 函数表示被保护的资源的调用函数
// resultWithBreaker 表示熔断器处于 open 状态的默认返回值
func wrapEntry(resource string, operate func() (interface{}, error), resultWithBreaker interface{}, errs ...error) (interface{}, error) {
	// 指定
	e, b := api.Entry(resource)
	if b != nil {
		// g1 blocked
		return resultWithBreaker, getError(b.BlockType().String())
	} else {
		result, err := operate()
		if err != nil {
			realErr := true
			for _, specifyErr := range errs {
				if specifyErr == err {
					realErr = false
				}
			}
			if realErr {
				api.TraceError(e, err)
			}
		}
		e.Exit()
		return result, err
	}
}

// wrapEntryWithoutResult 熔断装饰函数(无默认返回值)
// resource 代表资源名字
// operate 函数表示被保护的资源的调用函数
func wrapEntryWithoutResult(resource string, operate func() error, errs ...error) error {

	e, b := api.Entry(resource)
	if b != nil {
		// g1 blocked
		return getError(b.BlockType().String())
	} else {
		err := operate()
		if err != nil {
			realErr := true
			for _, specifyErr := range errs {
				if specifyErr == err {
					realErr = false
				}
			}
			if realErr {
				api.TraceError(e, err)
			}
		}
		e.Exit()
		return err
	}
}
