package user

import (
	"log"

	pb "github.com/gemdivk/LUMERA-SPA/user-service/proto"
	"google.golang.org/grpc"
)

var UserClient pb.UserServiceClient

func InitGRPCClient() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to user-service: %v", err)
	}
	UserClient = pb.NewUserServiceClient(conn)
}
