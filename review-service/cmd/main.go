package main

import (
	"database/sql"
	gr "github.com/gemdivk/LUMERA-SPA/review-service/internal/adapters/grpc"
	"github.com/gemdivk/LUMERA-SPA/review-service/internal/infrastructure/postgres"
	"github.com/gemdivk/LUMERA-SPA/review-service/internal/usecase"
	reviewpb "github.com/gemdivk/LUMERA-SPA/review-service/proto"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	repo := postgres.NewReviewRepo(db)
	uc := usecase.NewReviewInteractor(repo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	reviewpb.RegisterReviewServiceServer(server, gr.NewReviewServer(uc))

	log.Println("ReviewService gRPC listening on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
