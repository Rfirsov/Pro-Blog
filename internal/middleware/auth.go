package middleware

import (
	"net/http"
	"strings"
	"time"

	customErrors "github.com/Rfirsov/Pro-Blog/internal/errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/time/rate"
)

// AuthMiddleware verifies JWT tokens in incoming requests
func AuthMiddleware(jwtSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": customErrors.ErrAuthorizationHeaderMissing.Error()})
			c.Abort()
			return
		}

		// Check Bearer scheme
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": customErrors.ErrInvalidAuthorizationFormat.Error()})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Parse and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": customErrors.ErrInvalidTokenSignature.Error()})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": customErrors.ErrInvalidOrExpiredToken.Error()})
			}
			c.Abort()
			return
		}

		// Extract and validate claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": customErrors.ErrInvalidTokenClaims.Error()})
			c.Abort()
			return
		}

		// Check token expiration
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": customErrors.ErrTokenExpired.Error()})
				c.Abort()
				return
			}
		}

		// Set user information in context
		c.Set("user_id", claims["user_id"])
		c.Set("email", claims["email"])

		c.Next()
	}
}

// RateLimiter middleware to prevent brute force attacks
func RateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(time.Second), 10)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": customErrors.ErrTooManyRequests.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
