package controllers

import (
	"codewave/models"
	"codewave/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProjectRoutes(router *gin.RouterGroup) {
	router.POST("/projects", CreateProject)
	router.GET("/projects/:id", GetProject)
	router.GET("/projects", ListProjects)
	router.PUT("/projects/:id", UpdateProject)
}

func CreateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateProject(&project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, project)
}

func GetProject(c *gin.Context) {
	id := c.Param("id")
	project, err := services.GetProject(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}
	c.JSON(http.StatusOK, project)
}

func ListProjects(c *gin.Context) {
	// Convert user_id to uint
	userID, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return

	}
	projects, err := services.ListProjects(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
}

func UpdateProject(c *gin.Context) {
	id := c.Param("id")

	// Get project by id
	project, err := services.GetProject(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Bind JSON to project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update project
	if err := services.UpdateProject(project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}
