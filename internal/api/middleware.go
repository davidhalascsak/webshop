package api

import (
	"errors"
	"net/http"
	"strings"
	"webshop/internal/apperrors"
	"webshop/internal/auth"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		var appErr *apperrors.AppError
		if !errors.As(err, &appErr) {
			ErrorResponse(c, http.StatusInternalServerError, "internal_error", err)
			return
		}

		switch appErr.Kind {
		case apperrors.ErrBadRequest:
			ErrorResponse(c, http.StatusBadRequest, appErr.Message, appErr.Kind)
		case apperrors.ErrNotFound:
			ErrorResponse(c, http.StatusNotFound, appErr.Message, appErr.Kind)
		case apperrors.ErrConflict:
			ErrorResponse(c, http.StatusConflict, appErr.Message, appErr.Kind)
		case apperrors.ErrUnauthorized:
			ErrorResponse(c, http.StatusUnauthorized, appErr.Message, appErr.Kind)
		case apperrors.ErrForbidden:
			ErrorResponse(c, http.StatusForbidden, appErr.Message, appErr.Kind)
		case apperrors.ErrInternal:
			ErrorResponse(c, http.StatusInternalServerError, appErr.Message, appErr.Kind)
		default:
			ErrorResponse(c, http.StatusInternalServerError, "internal_error", err)
		}
	}
}

func AuthMiddleware(authenticator *auth.Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			ErrorResponse(c, http.StatusUnauthorized, "missing_token", apperrors.ErrUnauthorized)
			c.Abort()
			return
		}

		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(token, bearerPrefix) {
			ErrorResponse(c, http.StatusUnauthorized, "invalid_token", apperrors.ErrUnauthorized)
			c.Abort()
			return
		}

		user, err := authenticator.ValidateToken(c.Request.Context(), strings.TrimPrefix(token, bearerPrefix))
		if err != nil {
			ErrorResponse(c, http.StatusUnauthorized, "invalid_token", apperrors.ErrUnauthorized)
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func RequireRole(roles ...auth.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		userValue, ok := c.Get("user")
		user, validUser := userValue.(auth.User)
		if !ok || !validUser {
			ErrorResponse(c, http.StatusUnauthorized, "unauthorized", apperrors.ErrUnauthorized)
			c.Abort()
			return
		}

		for _, role := range roles {
			if user.HasRole(role) {
				c.Next()
				return
			}
		}

		ErrorResponse(c, http.StatusForbidden, "forbidden", apperrors.ErrForbidden)
		c.Abort()
	}
}
