package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
	"robot-go/service"
)

// ChatHandler 处理聊天请求
func ChatHandler(c *gin.Context) {
	var request struct {
		Question string `json:"question"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Invalid request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if request.Question == "" {
		log.Print("Empty question received")
		c.JSON(http.StatusBadRequest, gin.H{"error": "question cannot be empty"})
		return
	}

	ctx := context.Background()

	// 设置 SSE 头部
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Flush()

	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, request.Question),
	}

	// 调用流式 API
	_, err := service.LLMClient.GenerateContent(ctx, content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Fprintf(c.Writer, "data: %s\n\n", string(chunk))
		c.Writer.Flush()
		return nil
	}))

	if err != nil {
		log.Printf("Failed to generate content: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get response"})
		return
	}

	fmt.Fprintln(c.Writer, "data: [DONE]\n")
	c.Writer.Flush()
}
