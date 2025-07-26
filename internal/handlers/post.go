package handlers

import (
	"net/http"

	customErrors "github.com/Rfirsov/Pro-Blog/internal/errors"
	"github.com/Rfirsov/Pro-Blog/internal/models"
	"github.com/Rfirsov/Pro-Blog/internal/service"
	"github.com/Rfirsov/Pro-Blog/internal/utils"
	"github.com/gin-gonic/gin"
)

type PostHandler interface {
	CreatePost(c *gin.Context)
	UpdatePost(c *gin.Context)
	GetAllPosts(c *gin.Context)
	GetPostByID(c *gin.Context)
	DeletePost(c *gin.Context)
}

type postHandler struct {
	service service.PostService
}

func NewPostHandler(s service.PostService) *postHandler {
	return &postHandler{service: s}
}

// CreatePost godoc
// @Summary      Create a new blog post
// @Description  Create a post by an authenticated user
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        post  body      models.CreatePostRequest  true  "Post data"
// @Success      201   {object}  models.CreatePostSuccessResponse
// @Failure      400   {object}  models.CreatePostFailureBadRequestResponse
// @Failure      401   {object}  models.CreatePostFailureUnauthorizedResponse
// @Failure      500   {object}  models.CreatePostFailureInternalServerErrorResponse
// @Security     ApiKeyAuth
// @Router       /api/v1/posts [post]
func (h *postHandler) CreatePost(c *gin.Context) {
	var req models.CreatePostRequest
	userID, exists, err := utils.GetUUIDFromMiddleware(c, "user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": customErrors.ErrUserNotAuthenticated.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": customErrors.ErrUserNotAuthenticated.Error()})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": customErrors.ErrInvalidPostData.Error()})
		return
	}

	post, err := h.service.CreatePost(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": customErrors.ErrPostCreation.Error(), "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"post": post})
}

// UpdatePost godoc
// @Summary      Update an existing blog post
// @Description  Update a post by its ID
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Post ID (UUID)"
// @Param        post  body      models.UpdatePostRequest  true  "Post data"
// @Success      200   {object}  models.UpdatePostSuccessResponse
// @Failure      400   {object}  models.UpdatePostFailureBadRequestResponse
// @Failure      401   {object}  models.UpdatePostFailureUnauthorizedResponse
// @Failure      500   {object}  models.UpdatePostFailureInternalServerErrorResponse
// @Security     ApiKeyAuth
// @Router       /api/v1/posts/{id} [patch]
func (h *postHandler) UpdatePost(c *gin.Context) {
	var req models.UpdatePostRequest
	postID, err := utils.GetUUIDFromParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": customErrors.ErrPostId.Error()})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": customErrors.ErrInvalidPostData.Error()})
		return
	}

	updatedPost, err := h.service.UpdatePost(postID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": customErrors.ErrPostUpdate.Error(), "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post updated successfully", "post": updatedPost})
}

// GetAllPosts godoc
// @Summary      Get all posts
// @Description  Retrieve all blog posts
// @Tags         posts
// @Produce      json
// @Success      200   {object}  models.GetAllPostsSuccessResponse
// @Failure      500   {object}  models.GetAllPostsFailureInternalServerErrorResponse
// @Router       /api/v1/posts [get]
func (h *postHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.service.GetAllPosts()
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": customErrors.ErrPostNotFound.Error(), "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

// GetPostByID godoc
// @Summary      Get post by ID
// @Description  Retrieve a post by its UUID
// @Tags         posts
// @Produce      json
// @Param        id   path      string  true  "Post ID (UUID)"
// @Success      200  {object}  models.GetPostByIDSuccessResponse
// @Failure      400  {object}  models.GetPostByIDFailureBadRequestResponse
// @Failure      500  {object}  models.GetPostByIDFailureInternalServerErrorResponse
// @Router       /api/v1/posts/{id} [get]
func (h *postHandler) GetPostByID(c *gin.Context) {
	postID, err := utils.GetUUIDFromParam(c, "id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": customErrors.ErrPostNotFound.Error()})
		return
	}

	post, err := h.service.GetPostByID(postID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": customErrors.ErrPostNotFound.Error(), "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

// DeletePost godoc
// @Summary      Delete post
// @Description  Delete a blog post by its ID
// @Tags         posts
// @Produce      json
// @Param        id   path      string  true  "Post ID (UUID)"
// @Success      200  {object}  models.DeletePostSuccessResponse
// @Failure      400  {object}  models.DeletePostFailureBadRequestResponse
// @Failure      401  {object}  models.DeletePostFailureUnauthorizedResponse
// @Failure      500  {object}  models.DeletePostFailureInternalServerErrorResponse
// @Security     ApiKeyAuth
// @Router       /api/v1/posts/{id} [delete]
func (h *postHandler) DeletePost(c *gin.Context) {
	id, err := utils.GetUUIDFromParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": customErrors.ErrPostNotFound.Error()})
		return
	}

	err = h.service.DeletePost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": customErrors.ErrPostDelete.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "post deleted successfully",
		"post_id": id,
	})
}
