package controllers

import (
	"codewave/models"
	"codewave/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", Create)
		authRoutes.GET("/login", Login)
	}

	router.GET("/users/:id", GetUser)

}

func Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var loginReq models.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Print login request details
	fmt.Println("Login request - Email:", loginReq.Email, "Password:", loginReq.Password)

	token, err := services.AuthenticateUser(loginReq.Email, loginReq.Password)
	if err != nil {
		fmt.Println("Authentication error:", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
