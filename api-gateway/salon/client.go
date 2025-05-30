package salon

import (
	"log"

	pb "github.com/gemdivk/LUMERA-SPA/salon-management-service/proto"
	"google.golang.org/grpc"
)

var Client pb.SalonServiceClient

func InitGRPCClient() {
	conn, err := grpc.Dial("localhost:5053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to salon service: %v", err)
	}
	Client = pb.NewSalonServiceClient(conn)
}
