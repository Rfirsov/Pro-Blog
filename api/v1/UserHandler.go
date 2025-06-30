package apiV1

import (
	"net/http"

	"github.com/Rfirsov/Pro-Blog/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var UserHandler = &userHandler{}

type userHandler struct{}

func (h *userHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	uuidValue, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user := user.User{
		ID:    uuidValue,
		Name:  "Example User",
		Email: "somemails@gmail.com",
	}

	c.JSON(http.StatusOK, user)
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var user user.User

	user.ID = uuid.New()
	user.Name = "New User"
	user.Email = "somemails@gmail.com"

	c.JSON(http.StatusCreated, user)
}
