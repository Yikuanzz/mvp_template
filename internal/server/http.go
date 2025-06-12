package server

import (
	"fmt"
	"net/http"
	"time"

	"mvp/config"
	"mvp/internal/handler"
	"mvp/utils/middleware"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// ProviderSet 提供 HTTP 服务器和用户路由
var ProviderSet = wire.NewSet(
	NewHTTPServer,
)

// NewHTTPServer 创建 HTTP 服务器
func NewHTTPServer(userRouter *handler.UserHandler, conf *config.Config) *http.Server {
	router := gin.Default()
	gin.SetMode(conf.Server.Mode)
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(middleware.Cors())

	api := router.Group("/api/v1")
	apiPublic := api
	// apiAuth := api.Use(middleware.AuthMiddleware(conf.Auth.SecretKey))
	// 用户路由
	{
		apiPublic.POST("/user/login", userRouter.Login)
		apiPublic.POST("/user/register", userRouter.Register)
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(conf.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(conf.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(conf.Server.IdleTimeout) * time.Second,
	}

	return srv
}
