package routes

import (
	"github.com/0xMarvell/watchlist/pkg/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes handles all routing for user authentication
func RegisterUserRoutes(r *gin.Engine) {
	baseURl := r.Group("/api/v1")
	{
		baseURl.POST("/user/signup", controllers.Signup)
		baseURl.POST("/user/login", controllers.Login)
	}
}
