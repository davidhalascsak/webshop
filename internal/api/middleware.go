package api

import (
	"errors"
	"net/http"
	"webshop/internal/apperrors"

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
		case apperrors.ErrInternal:
			ErrorResponse(c, http.StatusInternalServerError, appErr.Message, appErr.Kind)
		default:
			ErrorResponse(c, http.StatusInternalServerError, "internal_error", err)
		}
	}
}
