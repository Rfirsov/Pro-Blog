package apiV1

import (
	"net/http"

	models "github.com/Rfirsov/Pro-Blog/internal/user"

	"github.com/gin-gonic/gin"
)

var UserHandler = &userHandler{}

type userHandler struct{}

func (h *userHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user := models.User{
		ID:    id,
		Name:  "Example User",
		Email: "somemails@gmail.com",
	}

	c.JSON(http.StatusOK, user)
}
