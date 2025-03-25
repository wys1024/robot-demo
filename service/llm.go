package service

import (
	"fmt"
	"github.com/tmc/langchaingo/llms/ollama"
)

// LLMClient 全局LLM客户端
var LLMClient *ollama.LLM

// InitLLM 初始化LLM客户端
func InitLLM(model string, apiHost string) error {
	var err error
	options := []ollama.Option{
		ollama.WithModel(model),
	}

	// 如果提供了API主机地址，则使用它
	if apiHost != "" {
		options = append(options, ollama.WithServerURL(apiHost))
	}

	LLMClient, err = ollama.New(options...)
	if err != nil {
		return fmt.Errorf("failed to initialize LLM: %v", err)
	}
	return nil
}
