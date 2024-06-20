package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GenRsUID() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.New().String())
		c.Next()
	}
}
