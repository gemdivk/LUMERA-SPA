package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"strings"

	"google.golang.org/grpc/metadata"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}

		// ✅ добавляем в gRPC metadata
		md := metadata.New(map[string]string{"authorization": auth})
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		// перезаписываем контекст в gin
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
