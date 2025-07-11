package customErrors

import "errors"

var (
	ErrAuthorizationHeaderMissing = errors.New("authorization header missing")
	ErrDatabase                   = errors.New("database error")
	ErrEmailAlreadyRegistered     = errors.New("email already registered")
	ErrInvalidAuthorizationFormat = errors.New("invalid authorization format")
	ErrInvalidCredentials         = errors.New("invalid credentials")
	ErrInvalidEmailFormat         = errors.New("invalid email format")
	ErrInvalidInputFormat         = errors.New("invalid input format")
	ErrInvalidLoginData           = errors.New("invalid login data")
	ErrInvalidOrExpiredToken      = errors.New("invalid or expired token")
	ErrInvalidTokenClaims         = errors.New("invalid token claims")
	ErrInvalidTokenSignature      = errors.New("invalid token signature")
	ErrLoginProcess               = errors.New("login process failed")
	ErrPasswordProcessing         = errors.New("password processing failed")
	ErrTokenExpired               = errors.New("token expired")
	ErrTokenGeneration            = errors.New("token generation failed")
	ErrTokenRefresh               = errors.New("token refresh failed")
	ErrTooManyRequests            = errors.New("too many requests")
	ErrTransactionCommit          = errors.New("transaction commit failed")
	ErrTransactionStart           = errors.New("transaction start failed")
	ErrUserCreation               = errors.New("user creation failed")
	ErrUserNotAuthenticated       = errors.New("user not authenticated")
	ErrUserNotFound               = errors.New("user not found")
)
