package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/0xMarvell/watchlist/pkg/config"
	"github.com/0xMarvell/watchlist/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Signup allows a user to register new account with the expected user details
func Signup(c *gin.Context) {
	// Get user data off request body and bind to payload struct
	var signupPayload struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if c.Bind(&signupPayload) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read request body",
		})
		return
	}
	// Hash password
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(signupPayload.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
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
		return
	}
	// Return JSON response to confirm operation success
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "new user successfully created",
	})
}

// Login allows existing user to login to their account
func Login(c *gin.Context) {
	// Get user data off request body and bind to payload struct
	var loginPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if c.Bind(&loginPayload) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read request body",
		})
		return
	}
	// Search database for user (using email as params)
	var user models.User
	config.DB.First(&user, "email = ?", loginPayload.Email)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found. invalid email or password",
		})
		return
	}
	// Compare password gotten from payload with password hash stored in database
	pwdErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginPayload.Password))
	if pwdErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}
	// Generate JWT token and store in httpOnly cookie
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject":         user.ID,
		"expiration_time": time.Now().Add(time.Hour * 2).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	var maxAge int = 3600 * 2
	var secure, httpOnly bool = false, true
	c.SetCookie("auth_token", tokenString, maxAge, "", "", secure, httpOnly)
	// return JSON response to confirm operation success
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "token generated successfully",
	})
}
