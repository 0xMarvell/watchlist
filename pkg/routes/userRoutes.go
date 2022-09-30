package routes

import (
	"github.com/0xMarvell/watchlist/pkg/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes handles all routing for user authentication
func RegisterUserRoutes(r *gin.Engine) {
	baseURl := r.Group("/api/v1/user")
	{
		baseURl.POST("/signup", controllers.Signup)
		baseURl.POST("/login", controllers.Login)
	}
}
