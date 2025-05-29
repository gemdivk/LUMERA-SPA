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

	grpcAdapter "github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/adapters/grpc"
	repoMongo "github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/infrastructure/mongo"
	"github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/usecase"
	pb "github.com/gemdivk/LUMERA-SPA/salon-management-service/proto"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	client, err := mongoDriver.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("lumera")

	repo := repoMongo.NewMongoRepo(db)
	uc := usecase.NewSalonUsecase(repo)

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSalonServiceServer(s, grpcAdapter.NewSalonServer(uc))

	log.Println("SalonManagementService gRPC started on :50053")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
