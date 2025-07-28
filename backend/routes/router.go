package routes

import (
	"github.com/Rfirsov/Pro-Blog/app"
	"github.com/Rfirsov/Pro-Blog/config"
	"github.com/Rfirsov/Pro-Blog/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	// Initialize router with middleware
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	postHandler := app.InitializePostService()
	postStatusHandler := app.InitializePostStatusService()

	// Public routes
	publicAuth := router.Group("/api/v1")
	{
		publicAuth.POST("/register", authHandler.Register)
		publicAuth.POST("/login", authHandler.Login)

		// Public post routes
		publicAuth.GET("/posts", postHandler.GetAllPosts)
		publicAuth.GET("/posts/:id", postHandler.GetPostByID)
	}

	// Protected routes with JWT middleware
	protectedAuth := router.Group("/api/v1")
	protectedAuth.Use(middleware.AuthMiddleware([]byte(config.Cfg.JWT.Secret)))
	{
		protectedAuth.POST("/refresh-token", authHandler.RefreshToken)
		protectedAuth.POST("/logout", authHandler.Logout)
		protectedAuth.GET("/profile", authHandler.GetUserProfile)
	}

	// Protected Post routes
	postRoutes := protectedAuth.Group("/posts")
	{
		postRoutes.POST("", postHandler.CreatePost)
		postRoutes.GET("/statuses", postStatusHandler.GetPostStatuses)
		postRoutes.PATCH("/:id", postHandler.UpdatePost)
		postRoutes.DELETE("/:id", postHandler.DeletePost)
	}

	return router
}
