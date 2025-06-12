package main

import (
	"flag"
	"net/http"
	"os"

	"mvp/config"
	"mvp/utils/app"
	"mvp/utils/log"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name 应用名称
	Name = "mvp"

	// Version 应用版本
	Version string

	// configPath 配置文件路径
	configPath string
)

func init() {
	flag.StringVar(&configPath, "conf", ".", "config path, eg: -conf config.toml")
}

// newApp 创建应用
func newApp(logger log.Logger, conf *config.Config, server *http.Server) *app.App {
	return app.New(
		app.Name(Name),
		app.Version(Version),
		app.Config(conf),
		app.Log(logger),
		app.Server(server),
	)
}

func main() {
	flag.Parse()

	// Logger 日志
	logger := log.NewLogger()

	// Config 配置
	conf, err := config.NewConfig(configPath)
	if err != nil {
		logger.Error("加载配置失败", "error", err)
		os.Exit(1)
	}

	// 使用 wire 生成的 initApp 实现
	app, cleanup, err := initApp(logger, conf)
	if err != nil {
		logger.Error("初始化应用失败", "error", err)
		os.Exit(1)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		logger.Error("运行应用失败", "error", err)
		os.Exit(1)
	}
}
