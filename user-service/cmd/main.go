package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	usergrpc "github.com/gemdivk/LUMERA-SPA/user-service/internal/adapters/grpc"
	"github.com/gemdivk/LUMERA-SPA/user-service/internal/infrastructure/postgres"
	"github.com/gemdivk/LUMERA-SPA/user-service/internal/usecase"
	pb "github.com/gemdivk/LUMERA-SPA/user-service/proto"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	repo := postgres.NewUserRepo(db)
	uc := usecase.NewUserInteractor(repo)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, usergrpc.NewUserServer(uc))

	log.Println("UserService gRPC running on port 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
