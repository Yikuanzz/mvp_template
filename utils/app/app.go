package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// App 应用
type App struct {
	options options
	ctx     context.Context
	cancel  context.CancelFunc
}

// New 创建应用
func New(opts ...Option) *App {
	o := options{
		ctx:  context.Background(),
		sigs: []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}

	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancel := context.WithCancel(o.ctx)

	return &App{
		options: o,
		ctx:     ctx,
		cancel:  cancel,
	}
}

// Name 返回应用名称
func (a *App) Name() string {
	return a.options.name
}

// Version 返回应用版本
func (a *App) Version() string {
	return a.options.version
}

// Run 运行应用
func (a *App) Run() error {
	// 启动前回调
	for _, fn := range a.options.beforeStart {
		if err := fn(a.ctx); err != nil {
			return err
		}
	}

	// 启动 HTTP 服务器
	if a.options.server != nil {
		go func() {
			_ = a.options.server.ListenAndServe()
		}()
	}

	// 启动后回调
	for _, fn := range a.options.afterStart {
		if err := fn(a.ctx); err != nil {
			return err
		}
	}

	// 监听信号或 context.Done
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, a.options.sigs...)
	select {
	case <-a.ctx.Done():
		// context 被取消
	case <-sigCh:
		// 收到退出信号
	}
	return a.Stop()
}

// Stop 停止应用
func (a *App) Stop() error {
	// 停止前回调
	for _, fn := range a.options.beforeStop {
		if err := fn(a.ctx); err != nil {
			return err
		}
	}

	// 优雅关闭 HTTP 服务器
	if a.options.server != nil {
		_ = a.options.server.Shutdown(a.ctx)
	}

	// 停止后回调
	for _, fn := range a.options.afterStop {
		if err := fn(a.ctx); err != nil {
			return err
		}
	}

	a.cancel()
	return nil
}
