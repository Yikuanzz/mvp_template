package app

import (
	"context"
	"net/http"
	"os"

	"mvp/config"
	"mvp/utils/log"
)

// Option 应用选项
type Option func(o *options)

// options 应用选项
type options struct {
	name    string // 应用名称
	version string // 应用版本

	ctx  context.Context // 应用上下文
	sigs []os.Signal     // 应用信号

	conf   *config.Config // 应用配置
	log    log.Logger     // 应用日志
	server *http.Server   // 应用服务器

	beforeStart []func(context.Context) error // 应用启动前回调
	beforeStop  []func(context.Context) error // 应用停止前回调
	afterStart  []func(context.Context) error // 应用启动后回调
	afterStop   []func(context.Context) error // 应用停止后回调
}

// Name 设置应用名称
func Name(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

// Version 设置应用版本
func Version(version string) Option {
	return func(o *options) {
		o.version = version
	}
}

// Context 设置应用上下文
func Context(ctx context.Context) Option {
	return func(o *options) {
		o.ctx = ctx
	}
}

// Signals 设置应用信号
func Signals(sigs ...os.Signal) Option {
	return func(o *options) {
		o.sigs = sigs
	}
}

// Config 设置应用配置
func Config(conf *config.Config) Option {
	return func(o *options) {
		o.conf = conf
	}
}

// Log 设置应用日志
func Log(log log.Logger) Option {
	return func(o *options) {
		o.log = log
	}
}

// Server 设置应用服务器
func Server(server *http.Server) Option {
	return func(o *options) {
		o.server = server
	}
}

// BeforeStart 设置应用启动前回调
func BeforeStart(fn func(context.Context) error) Option {
	return func(o *options) {
		o.beforeStart = append(o.beforeStart, fn)
	}
}

// BeforeStop 设置应用停止前回调
func BeforeStop(fn func(context.Context) error) Option {
	return func(o *options) {
		o.beforeStop = append(o.beforeStop, fn)
	}
}

// AfterStart 设置应用启动后回调
func AfterStart(fn func(context.Context) error) Option {
	return func(o *options) {
		o.afterStart = append(o.afterStart, fn)
	}
}

// AfterStop 设置应用停止后回调
func AfterStop(fn func(context.Context) error) Option {
	return func(o *options) {
		o.afterStop = append(o.afterStop, fn)
	}
}
