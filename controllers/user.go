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
		authRoutes.POST("/login", Login)
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

	token, err := services.AuthenticateUser(loginReq.Email, loginReq.Password)
	if err != nil {
		fmt.Println("Authentication error:", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	auth := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   3600,
		Secure:   false,
		HttpOnly: true,
	}

	c.SetCookie(auth.Name, auth.Value, auth.MaxAge, auth.Path, auth.Domain, auth.Secure, auth.HttpOnly)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
