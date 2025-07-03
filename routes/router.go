package routes

import (
	"github.com/Rfirsov/Pro-Blog/app"
	"github.com/Rfirsov/Pro-Blog/internal/user"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	_, _, userHandler := app.InitializeUserService()

	v1Group := router.Group("/api/v1")
	registerUserRoutes(v1Group, userHandler)

	return router
}

func registerUserRoutes(g *gin.RouterGroup, userHandler user.Handler) {
	users := g.Group("/users")
	users.POST("", userHandler.CreateUser)
	users.GET("/:id", userHandler.GetUser)
}
