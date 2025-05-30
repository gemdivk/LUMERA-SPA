package booking

import (
	"log"

	pb "github.com/gemdivk/LUMERA-SPA/booking-service/proto"
	"google.golang.org/grpc"
)

var BookingClient pb.BookingServiceClient

func InitGRPCClient() {
	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to booking-service: %v", err)
	}
	BookingClient = pb.NewBookingServiceClient(conn)
}
