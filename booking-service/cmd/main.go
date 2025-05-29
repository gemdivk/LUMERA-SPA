package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	grpcAdapter "github.com/gemdivk/LUMERA-SPA/booking-service/internal/adapters/grpc"
	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/infrastructure/mongo"
	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/usecase"
	pb "github.com/gemdivk/LUMERA-SPA/booking-service/proto"
)

func main() {
	_ = godotenv.Load()

	client, err := mongoDriver.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("lumera")
	repo := mongo.NewMongoRepo(db)
	uc := usecase.NewBookingUsecase(repo)

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterBookingServiceServer(s, grpcAdapter.NewBookingServer(uc))

	log.Println("BookingService running on :50054")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
