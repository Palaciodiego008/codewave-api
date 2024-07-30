package controllers

import (
	"codewave/models"
	"codewave/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpenAPI(c *gin.Context) {
	var openAPI models.OpenAPI
	if err := c.ShouldBindJSON(&openAPI); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "data": gin.H{"error": err.Error()}})
		return
	}

	err := services.CreateOpenAPI(&openAPI)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "data": gin.H{"error": err.Error()}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": openAPI})
}

func GetOpenAPI(c *gin.Context) {
	id := c.Param("id")
	openAPI, err := services.GetOpenAPIByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "data": gin.H{"error": err.Error()}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": openAPI})
}
