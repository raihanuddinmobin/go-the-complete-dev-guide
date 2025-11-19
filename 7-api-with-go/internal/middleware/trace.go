package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := uuid.New().String()

		ctx := context.WithValue(c.Request.Context(), "traceID", traceID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
