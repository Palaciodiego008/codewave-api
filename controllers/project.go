package controllers

import (
	"codewave/models"
	"codewave/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProjectRoutes(router *gin.Engine) {
	router.POST("/projects", CreateProject)
	router.GET("/projects/:id", GetProject)
	router.GET("/projects", ListProjects)
}

func CreateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obtener el ID del usuario desde el token
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	project.UserID = userID.(uint)

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
	projects, err := services.ListProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
}
