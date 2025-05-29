package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	usergrpc "github.com/gemdivk/LUMERA-SPA/user-service/internal/adapters/grpc"
	"github.com/gemdivk/LUMERA-SPA/user-service/internal/infrastructure/cache"
	"github.com/gemdivk/LUMERA-SPA/user-service/internal/infrastructure/postgres"
	"github.com/gemdivk/LUMERA-SPA/user-service/internal/usecase"
	pb "github.com/gemdivk/LUMERA-SPA/user-service/proto"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load(".env")
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

	nc, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}
	defer nc.Close()

	repo := postgres.NewUserRepo(db)
	allUsers, err := repo.GetAll()
	if err != nil {
		log.Fatalf("Failed to fetch users from DB: %v", err)
	}
	userCache := cache.NewUserCache()
	userCache.LoadInitial(allUsers)
	log.Printf("User cache initialized with %d users\n", len(allUsers))

	uc := usecase.NewUserInteractorWithCache(repo, nc, userCache)

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
