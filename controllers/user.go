package controllers

import (
	"codewave/models"
	"codewave/services"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	oauth2Config = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}
	oauth2State = "random"
)

func UserRoutes(router *gin.Engine) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", Create)
		authRoutes.POST("/login", Login)
		authRoutes.POST("/logout", Logout)
		authRoutes.GET("/github", GitHubLogin)
		authRoutes.GET("/github/callback", GitHubCallback)
	}

	router.GET("/users/:id", GetUser)
}

func Login(c *gin.Context) {
	var loginReq models.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "data": gin.H{"error": err.Error()}})
		return
	}

	token, err := services.AuthenticateUser(loginReq.Email, loginReq.Password)
	if err != nil {
		fmt.Println("Authentication error:", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "data": gin.H{"error": "Invalid email or password"}})
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost:3000", false, true)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"token": token}})
}

func Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "data": gin.H{"error": err.Error()}})
		return
	}
	err := services.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "data": gin.H{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": user})
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

func Logout(c *gin.Context) {
	auth := http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, &auth)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func GitHubLogin(c *gin.Context) {
	url := oauth2Config.AuthCodeURL(oauth2State)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GitHubCallback(c *gin.Context) {
	state := c.Query("state")
	if state != oauth2State {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state"})
		return
	}

	code := c.Query("code")
	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}

	client := oauth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
		return
	}

	c.JSON(http.StatusOK, userInfo)
}
