package main

import (
	gr "github.com/gemdivk/LUMERA-SPA/payment-service/internal/adapters/grpc"
	infrastructure "github.com/gemdivk/LUMERA-SPA/payment-service/internal/insfrastructure"
	repository "github.com/gemdivk/LUMERA-SPA/payment-service/internal/insfrastructure/postgres"
	"github.com/gemdivk/LUMERA-SPA/payment-service/internal/usecase"
	pb "github.com/gemdivk/LUMERA-SPA/payment-service/proto"
	"github.com/nats-io/nats.go"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}

	db := infrastructure.InitDB()
	repo := repository.New(db) // репозиторий и транзакционный раннер

	stripeClient := infrastructure.InitStripeClient()
	natsClient := infrastructure.NewNATSClient(nc)

	uc := usecase.New(repo, repo, stripeClient, natsClient)
	lis, _ := net.Listen("tcp", ":50055")
	server := grpc.NewServer()
	pb.RegisterPaymentServiceServer(server, &gr.Server{UC: uc})
	server.Serve(lis)
}
