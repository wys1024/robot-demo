package service

import (
	"fmt"
	"github.com/tmc/langchaingo/llms/ollama"
)

// LLMClient 全局LLM客户端
var LLMClient *ollama.LLM

// InitLLM 初始化LLM客户端
func InitLLM() error {
	var err error
	LLMClient, err = ollama.New(ollama.WithModel("qwen2:7b"))
	if err != nil {
		return fmt.Errorf("failed to initialize LLM: %v", err)
	}
	return nil
}
