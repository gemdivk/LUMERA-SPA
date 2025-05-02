package review

import (
	"log"

	pb "github.com/gemdivk/LUMERA-SPA/review-service/proto"
	"google.golang.org/grpc"
)

var ReviewClient pb.ReviewServiceClient

func InitGRPCClient() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // или TLS
	if err != nil {
		log.Fatalf("Не удалось подключиться к gRPC: %v", err)
	}
	ReviewClient = pb.NewReviewServiceClient(conn)
}
