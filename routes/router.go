package routes

import (
	"github.com/Rfirsov/Pro-Blog/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	{
		v1Group := router.Group("/api/v1")
		v1Group.GET("/users", apiV1.UserHandler.GetUser)
	}

	return router
}
