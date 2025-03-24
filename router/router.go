package router

import (
	"github.com/gin-gonic/gin"
	"robot-demo/handler"
)

// SetupRouter 初始化路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 添加静态文件服务
	r.Static("/static", "./static")

	// 首页路由
	r.GET("/", func(c *gin.Context) {
		c.File("static/index.html")
	})

	// chat路由
	r.POST("/api/chat", handler.ChatHandler)

	return r
}
