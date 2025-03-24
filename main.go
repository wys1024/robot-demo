package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"robot-go/router"
	"robot-go/service"
)

func main() {
	// 初始化LLM客户端
	if err := service.InitLLM(); err != nil {
		log.Fatalf("Failed to initialize LLM: %v", err)
	}

	// 设置gin模式
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化路由
	r := router.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9527"
	}

	log.Printf("Server is running on http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
