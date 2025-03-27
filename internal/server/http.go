package server

import (
	v1 "robot-demo/api/aichat/v1"
	"robot-demo/internal/conf"
	"robot-demo/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, aichat *service.AichatService, logger log.Logger) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)

	// TODO: 静态资源服务

	//1. 注册静态资源服务
	//注册静态资源服务，将静态资源文件映射到指定的URL路径上

	v1.RegisterAichatHTTPServer(srv, aichat)

	return srv
}
