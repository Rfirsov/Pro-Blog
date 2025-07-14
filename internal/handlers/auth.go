package handlers

import (
	"net/http"
	"time"

	customErrors "github.com/Rfirsov/Pro-Blog/internal/errors"
	"github.com/Rfirsov/Pro-Blog/internal/models"
	"github.com/Rfirsov/Pro-Blog/internal/service"
	"github.com/Rfirsov/Pro-Blog/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
	Logout(c *gin.Context)
	GetUserProfile(c *gin.Context)
}

type authHandler struct {
	service         service.AuthService
	tokenExpiration time.Duration
}

func NewAuthHandler(s service.AuthService, tokenExpiration time.Duration) *authHandler {
	return &authHandler{service: s, tokenExpiration: tokenExpiration}
}

// Register godoc
// @Summary      Register new user
// @Description  Register a new user with email and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.UserRegister  true  "User registration data"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /api/v1/register [post]
func (h *authHandler) Register(c *gin.Context) {
	var user models.UserRegisterRequest

	// Validate input JSON
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   customErrors.ErrInvalidInputFormat.Error(),
			"details": err.Error(),
		})
		return
	}

	// Additional validation
	if err := h.service.ValidateEmailFormat(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	exists, err := h.service.IfUserExists(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": customErrors.ErrDatabase.Error()})
		return
	}
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": customErrors.ErrEmailAlreadyRegistered.Error()})
		return
	}

	// Hash password
	hashedPassword, errHashedPassword := utils.HashPassword(user.Password)
	newUser := models.User{Email: user.Email, Password: hashedPassword}
	if errHashedPassword != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": customErrors.ErrPasswordProcessing.Error()})
		return
	}

	errRegistration := h.service.Registration(&newUser)
	if errRegistration != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errRegistration.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user_id": newUser.ID,
	})
}

// Login godoc
// @Summary      User login
// @Description  Authenticate user and get JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200          {object}  map[string]string
// @Failure      400          {object}  map[string]string
// @Failure      401          {object}  map[string]string
// @Failure      500          {object}  map[string]string
// @Router       /api/v1/login [post]
// Login handles user authentication and JWT generation
func (h *authHandler) Login(c *gin.Context) {
	var req models.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": customErrors.ErrInvalidLoginData.Error()})
		return
	}

	user, err := h.service.GetUser(req)

	if err == gorm.ErrRecordNotFound {
		// Don't specify whether email or password was wrong
		c.JSON(http.StatusUnauthorized, gin.H{"error": customErrors.ErrInvalidCredentials.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": customErrors.ErrLoginProcess.Error()})
		return
	}

	// Verify password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		// Use same message as above for security
		c.JSON(http.StatusUnauthorized, gin.H{"error": customErrors.ErrInvalidCredentials.Error()})
		return
	}

	// Generate JWT
	tokenString, err := h.service.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": customErrors.ErrTokenGeneration.Error()})
		return
	}

	// Return token with expiration
	c.JSON(http.StatusOK, gin.H{
		"token":      tokenString,
		"expires_in": h.tokenExpiration.Seconds(),
		"token_type": "Bearer",
	})
}

// RefreshToken godoc
// @Summary      Refresh JWT token
// @Description  Generate a new JWT token using a valid refresh token
// @Tags         auth
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/v1/refresh-token [post]
// @Security     ApiKeyAuth
// RefreshToken generates a new token for valid users
func (h *authHandler) RefreshToken(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": customErrors.ErrUserNotAuthenticated.Error()})
		return
	}

	tokenString, err := h.service.GenerateRefreshJWT(userID.(int))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": customErrors.ErrTokenRefresh.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":      tokenString,
		"expires_in": h.tokenExpiration.Seconds(),
		"token_type": "Bearer",
	})
}

// Logout godoc
// @Summary      Logout user
// @Description  Logout user and invalidate tokens on client side
// @Tags         auth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /api/v1/logout [post]
// @Security ApiKeyAuth
func (h *authHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":      "Successfully logged out",
		"instructions": "Please remove the token from your client storage",
	})
}

// GetUserProfile godoc
// @Summary      Get user profile
// @Description  Retrieve logged-in user's profile information
// @Tags         user
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]string
// @Security     ApiKeyAuth
// @Router       /api/v1/profile [get]
func (h *authHandler) GetUserProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	email, _ := c.Get("email")
	userRole, _ := c.Get("role")

	c.JSON(200, gin.H{
		"user_id": userID,
		"email":   email,
		"role":    userRole,
	})
}
