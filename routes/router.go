package routes

import (
	"github.com/Rfirsov/Pro-Blog/app"
	"github.com/Rfirsov/Pro-Blog/config"
	"github.com/Rfirsov/Pro-Blog/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// Initialize router with middleware
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Initialize handlers with JWT configuration
	authHandler := app.InitializeAuthService()

	// Public routes
	public := router.Group("/api/v1")
	{
		public.POST("/register", authHandler.Register)
		public.POST("/login", authHandler.Login)
	}

	// Protected routes with JWT middleware
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware([]byte(config.Cfg.JWT.Secret)))
	{
		protected.POST("/refresh-token", authHandler.RefreshToken)
		protected.POST("/logout", authHandler.Logout)
		protected.GET("/profile", authHandler.GetUserProfile)
	}

	return router
}

