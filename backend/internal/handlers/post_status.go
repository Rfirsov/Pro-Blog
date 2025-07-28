package handlers

import (
	"net/http"

	customErrors "github.com/Rfirsov/Pro-Blog/internal/errors"
	"github.com/Rfirsov/Pro-Blog/internal/service"
	"github.com/gin-gonic/gin"
)

type PostStatusHandler interface {
	GetPostStatuses(c *gin.Context)
}

type postStatusHandler struct {
	service service.PostStatusService
}

func NewPostStatusHandler(s service.PostStatusService) *postStatusHandler {
	return &postStatusHandler{service: s}
}

// GetPostStatuses godoc
// @Summary      List available post statuses
// @Description  Returns all possible statuses a post may have (draft, published, archived, etc.)
// @Tags         posts
// @Accept       json
// @Produce      json
// @Success      200 {object} models.GetPostStatusesSuccessResponse
// @Failure      500 {object} models.GetPostStatusesFailureInternalServerErrorResponse
// @Security     ApiKeyAuth
// @Router       /api/v1/posts/statuses [get]
func (h *postStatusHandler) GetPostStatuses(c *gin.Context) {
	statuses, err := h.service.GetPostStatuses()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": customErrors.ErrFetchPostStatuses.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statuses": statuses})
}
