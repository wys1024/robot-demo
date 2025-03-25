package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config 应用配置结构
type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Mode string `yaml:"mode"`
	} `yaml:"server"`
	LLM struct {
		Model   string `yaml:"model"`
		APIHost string `yaml:"api_host"`
	} `yaml:"llm"`
}

// GlobalConfig 全局配置实例
var GlobalConfig Config

// LoadConfig 从文件加载配置
func LoadConfig(path string) error {
	// 确保配置文件存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 如果配置文件不存在，创建默认配置
		return createDefaultConfig(path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, &GlobalConfig)
}

// createDefaultConfig 创建默认配置文件
func createDefaultConfig(path string) error {
	// 确保目录存在
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// 默认配置
	GlobalConfig = Config{
		Server: struct {
			Port string `yaml:"port"`
			Mode string `yaml:"mode"`
		}{
			Port: "9527",
			Mode: "debug",
		},
		LLM: struct {
			Model   string `yaml:"model"`
			APIHost string `yaml:"api_host"`
		}{
			Model:   "qwen2:7b",
			APIHost: "http://localhost:11434",
		},
	}

	// 序列化配置
	data, err := yaml.Marshal(GlobalConfig)
	if err != nil {
		return err
	}

	// 写入文件
	return os.WriteFile(path, data, 0644)
}
