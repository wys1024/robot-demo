package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"robot-demo/config"
	"robot-demo/router"
	"robot-demo/service"
)

func main() {
	// 加载配置文件
	configPath := filepath.Join("config", "config.yaml")
	if err := config.LoadConfig(configPath); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Printf("Config loaded successfully")

	// 初始化LLM客户端
	if err := service.InitLLM(config.GlobalConfig.LLM.Model, config.GlobalConfig.LLM.APIHost); err != nil {
		log.Fatalf("Failed to initialize LLM: %v", err)
	}

	// 设置gin模式
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = config.GlobalConfig.Server.Mode
	}
	gin.SetMode(ginMode)

	// 初始化路由
	r := router.SetupRouter()

	// 获取端口
	port := os.Getenv("PORT")
	if port == "" {
		port = config.GlobalConfig.Server.Port
	}

	log.Printf("Server is running on http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
