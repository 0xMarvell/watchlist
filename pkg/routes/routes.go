package routes

import (
	"github.com/0xMarvell/watchlist/pkg/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes handles all routing for user authentication
func RegisterRoutes(r *gin.Engine) {
	// User Routes
	userRoutes := r.Group("/api/v1/user")
	{
		userRoutes.POST("/signup", controllers.Signup)
		userRoutes.POST("/login", controllers.Login)
	}

	// Watchlist Routes
	watchlistRoutes := r.Group("api/v1/watchlist")
	{
		watchlistRoutes.GET("", controllers.AddShow)
	}
}
