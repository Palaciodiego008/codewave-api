package controllers

import (
	"codewave/models"
	"codewave/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OpenAPIRoutes(router *gin.RouterGroup) {
	router.POST("/openapis", CreateOpenAPI)
	router.GET("/openapis/:id", GetOpenAPI)
	router.GET("/openapis", ListOpenApis)
}

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

func ListOpenApis(c *gin.Context) {
	// Convert user_id to uint
	userID, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return

	}

	openAPIs, err := services.ListOpenAPIs(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "data": gin.H{"error": err.Error()}})
		return
	}

	fmt.Println("openAPIs", openAPIs)

	c.JSON(http.StatusOK, gin.H{"success": true, "data": openAPIs})

}
