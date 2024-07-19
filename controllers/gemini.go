package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GeminiRoutes(router *gin.RouterGroup) {
	router.POST("/chat-gemini", ChatGeminiResponse)
}

func ChatGeminiResponse(c *gin.Context) {
	var geminiRequest struct {
		Prompt string `json:"prompt"`
	}
	if err := c.ShouldBindJSON(&geminiRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	google_api_key := os.Getenv("GOOGLE_API_KEY")
	if google_api_key == "" {
		google_api_key = "AIzaSyDMQ9txYz2_efafEkKVs6mEKvoi0YcTAo4"
	}
	client, err := genai.NewClient(ctx, option.WithAPIKey(google_api_key))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(geminiRequest.Prompt))

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"response": resp.Candidates[0].Content.Parts[0]})
}
