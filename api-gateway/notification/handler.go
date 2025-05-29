package notification

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func GetLogs(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
		return
	}

	md := metadata.New(map[string]string{"authorization": token})
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	resp, err := NotificationClient.GetEmailLogs(ctx, &emptypb.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
