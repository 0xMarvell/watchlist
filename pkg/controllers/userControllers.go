package controllers

import (
	"net/http"

	"github.com/0xMarvell/watchlist/pkg/config"
	"github.com/0xMarvell/watchlist/pkg/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Signup allows a user to register new account with the expected user details
func Signup(c *gin.Context) {
	// Get data off req body
	var signupPayload struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// Hash password
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(signupPayload.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
	}
	// Store details in DB
	user := models.User{
		Name:     signupPayload.Name,
		Email:    signupPayload.Email,
		Password: string(pwdHash),
	}
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
	}
	// Return JSON response to confirm operation success
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "new user successfully created",
	})
}
