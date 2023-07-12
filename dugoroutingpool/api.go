package dugoroutingpool

import (
	"context"
	"github.com/Jeffail/tunny"
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

type GoRoutingPool interface {
	Close()
	Tune(size int)
}

// GoRoutingAsyncPool 异步协程池
type GoRoutingAsyncPool struct {
	Pool       *ants.PoolWithFunc
	Wg         sync.WaitGroup
	GLock      sync.RWMutex
	Size       int
	HandleFunc func(line interface{})
}

// GoRoutingSyncPool 同步协程池
type GoRoutingSyncPool struct {
	Pool       *tunny.Pool
	Size       int
	HandleFunc func(line interface{}) interface{}
}

// NewAsyncPool 异步协程池 size池的大小  f 自定义处理函数  options 配置
func NewAsyncPool(size int, f func(line interface{}), options ...ants.Option) *GoRoutingAsyncPool {
	handle := &GoRoutingAsyncPool{
		Size:       size,
		HandleFunc: f,
	}
	handle.Pool, _ = ants.NewPoolWithFunc(handle.Size, func(line interface{}) {
		handle.HandleFunc(line)
		handle.Wg.Done()
	}, options...)
	return handle
}

// NewSyncPool 同步协程池 size池的大小  f 自定义处理函数
func NewSyncPool(size int, f func(line interface{}) interface{}) *GoRoutingSyncPool {
	handle := &GoRoutingSyncPool{
		Size:       size,
		HandleFunc: f,
	}
	handle.Pool = tunny.NewFunc(handle.Size, f)
	return handle
}

// Process 处理数据
func (syncP *GoRoutingSyncPool) Process(data interface{}) interface{} {
	return syncP.Pool.Process(data)
}

// ProcessTimed 处理数据 支持超时
func (syncP *GoRoutingSyncPool) ProcessTimed(data interface{}, timeout time.Duration) (interface{}, error) {
	return syncP.Pool.ProcessTimed(data, timeout)
}

// ProcessCtx 处理数据 支持上下文
func (syncP *GoRoutingSyncPool) ProcessCtx(ctx context.Context, data interface{}) (interface{}, error) {
	return syncP.Pool.ProcessCtx(ctx, data)
}

// Tune 动态调整池的容量
func (syncP *GoRoutingSyncPool) Tune(resize int) {
	syncP.Pool.SetSize(resize)
}

// Close 释放资源
func (syncP *GoRoutingSyncPool) Close() {
	if syncP.Pool != nil {
		syncP.Pool.Close()
	}
}

// SendLine 发送数据
func (asyncP *GoRoutingAsyncPool) SendLine(line interface{}) error {
	asyncP.Wg.Add(1)
	err := asyncP.Pool.Invoke(line)
	if err != nil {
		return err
	}
	return nil
}

// Wait 等待所有任务完成
func (asyncP *GoRoutingAsyncPool) Wait() {
	asyncP.Wg.Wait()
}

// Done 任务完成
func (asyncP *GoRoutingAsyncPool) Done() {
	asyncP.Wg.Done()
}

// Add 添加任务
func (asyncP *GoRoutingAsyncPool) Add(n int) {
	asyncP.Wg.Add(n)
}

// Close 释放资源
func (asyncP *GoRoutingAsyncPool) Close() {
	if asyncP.Pool != nil {
		asyncP.Pool.Release()
	}
}

//	Lock 加锁
func (asyncP *GoRoutingAsyncPool) Lock() {
	asyncP.GLock.Lock()
}

// UnLock 解锁
func (asyncP *GoRoutingAsyncPool) UnLock() {
	asyncP.GLock.Unlock()
}

// Tune 动态调整池的容量
func (asyncP *GoRoutingAsyncPool) Tune(resize int) {
	asyncP.Pool.Tune(resize)
}

// SetClearNoUsedGoroutine 设置定时清理无用goroutines
func SetClearNoUsedGoroutine(interval time.Duration) ants.Option {
	return ants.WithExpiryDuration(interval)
}

// SetPreAlloc 设置预分配内存
func SetPreAlloc(preAlloc bool) ants.Option {
	return ants.WithPreAlloc(preAlloc)
}

// SetPanicHandler 设置异常处理
func SetPanicHandler(panicHandler func(interface{})) ants.Option {
	return ants.WithPanicHandler(panicHandler)
}

// SetNonblocking 设置池是否阻塞 默认阻塞
func SetNonblocking(nonblocking bool) ants.Option {
	return ants.WithNonblocking(nonblocking)
}

// SetLogger 设置logger
func SetLogger(logger ants.Logger) ants.Option {
	return ants.WithLogger(logger)
}

// SetMaxBlockingTasks 设置最大阻塞任务数量
func SetMaxBlockingTasks(maxBlockingTasks int) ants.Option {
	return ants.WithMaxBlockingTasks(maxBlockingTasks)
}

// IsSyncJobTimedOutError 判断是否超时
func IsSyncJobTimedOutError(err error) bool {
	//return errors.Is(err, tunny.ErrJobTimedOut)
	return err == tunny.ErrJobTimedOut
}

// IsSyncDeadlineExceededError 判断是否过期
func IsSyncDeadlineExceededError(err error) bool {
	return err == context.DeadlineExceeded
}
