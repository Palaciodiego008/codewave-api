package controllers

import (
	"codewave/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

func OpenAIRoutes(router *gin.RouterGroup) {
	router.POST("/chat-gpt", ChatGPTResponse)
}

func ChatGPTResponse(c *gin.Context) {
	ctx := c.Request.Context()
	var request models.OpenAIPrompt

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	openaiKey := os.Getenv("OPENAI_KEY")
	if openaiKey == "" {
		openaiKey = "sk-proj-171WNk8MLQW2Wy6liqzGT3BlbkFJ2P7iGNuo1HWpjP3CSr76"
	}

	client := openai.NewClient(openaiKey)
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 100,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: request.Prompt,
			},
		},
	}

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": resp.Choices[0].Message.Content})
}
