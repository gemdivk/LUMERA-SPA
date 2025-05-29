package notification

import (
	"log"

	pb "github.com/gemdivk/LUMERA-SPA/notification-service/proto"
	"google.golang.org/grpc"
)

var NotificationClient pb.NotificationServiceClient

func InitGRPCClient() {
	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to notification-service: %v", err)
	}
	NotificationClient = pb.NewNotificationServiceClient(conn)
}
