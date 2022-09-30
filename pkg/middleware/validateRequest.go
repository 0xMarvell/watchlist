package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/0xMarvell/watchlist/pkg/config"
	"github.com/0xMarvell/watchlist/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// ValidateRequest acts as middleware to validate the received JWT token
// before granting authorization to a user
func ValidateRequest(c *gin.Context) {
	// Get the cookie off request
	tokenString, cookieErr := c.Cookie("auth_token")
	if cookieErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "failed to get cookie off request",
		})
	}
	// Decode/Validate token stored in cookie
	token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if tokenErr != nil {
		log.Println("Failed to parse JWT token: ", tokenErr)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if token has not surpassed its expiry time
		if float64(time.Now().Unix()) > claims["expiration_time"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token is expired",
			})
		}
		// Find user using token subject
		var user models.User
		config.DB.First(&user, claims["subject"])
		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprintf("user with id %v does not exist", user.ID),
			})
		}
		// Attach user to request
		c.Set("user", user)
		// Authorize user and continue
		c.Next() // sends request from middleware to expected controller
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to validate JWT token",
		})
	}
}
