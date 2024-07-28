package controllers

import (
	"codewave/utils"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func AnalysisRoutes(router *gin.RouterGroup) {
	router.POST("/recommendation-analysis", RecommendationAnalysis)
}

func RecommendationAnalysis(c *gin.Context) {
	ctx := c.Request.Context()

	var geminiRequest struct {
		SnapshotCode string   `json:"snapshot_code"`
		Sections     []string `json:"sections"`
	}
	if err := c.ShouldBindJSON(&geminiRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	api_key := os.Getenv("GEMINI_API_KEY")
	if api_key == "" {
		api_key = "AIzaSyDtaynem1JgJyUELQeQqYKXpprJBWj1chg"
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(api_key))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	sections := []string{"Security", "Readability", "Static Code Analysis", "Dependency Scanning"}
	selectedSections := ""

	for _, section := range sections {
		if utils.Contains(geminiRequest.Sections, section) {
			selectedSections += section + ", "
		}
	}

	if len(selectedSections) > 0 {
		selectedSections = selectedSections[:len(selectedSections)-2]
	}

	prompt := `Please analyze the following code and provide a JSON response. 
	The JSON should have sections including ` + selectedSections + `. 
	Each section should contain items with a title, description (including security values or percentages), and status. 
	The status should be 'Passed', 'Needs Improvement', or 'Failed'. Only return the JSON, no more explanation for your part.
	
	Code:
	` + geminiRequest.SnapshotCode

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var geminiResponseContent string
	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		geminiResponseContent = fmt.Sprintf("%s", resp.Candidates[0].Content)
	} else {
		c.JSON(500, gin.H{"error": "No content received from Gemini"})
		return
	}
	fmt.Sprintln("Gemini Response Content: ", geminiResponseContent)

	geminiResponse, err := utils.ExtractJSON(geminiResponseContent)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"response": geminiResponse})
}
