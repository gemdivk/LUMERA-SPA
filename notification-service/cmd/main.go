package main

import (
	"database/sql"
	pb "github.com/gemdivk/LUMERA-SPA/notification-service/proto"
	"log"
	"net"
	"os"

	grpcAdapter "github.com/gemdivk/LUMERA-SPA/notification-service/internal/adapters/grpc"
	natsAdapter "github.com/gemdivk/LUMERA-SPA/notification-service/internal/adapters/nats"
	"github.com/gemdivk/LUMERA-SPA/notification-service/internal/infrastructure/mail"
	"github.com/gemdivk/LUMERA-SPA/notification-service/internal/infrastructure/postgres"
	"github.com/gemdivk/LUMERA-SPA/notification-service/internal/usecase"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	_ = godotenv.Load("./.env")

	connStr := "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	nc, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}
	defer nc.Close()

	repo := postgres.NewEmailLogRepo(db)
	smtpSender := mail.NewSmtpSender()
	uc := usecase.NewEmailUsecase(smtpSender, repo)

	natsAdapter.NewConsumer(nc, uc).Subscribe()

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServiceServer(grpcServer, grpcAdapter.NewNotificationServer(uc))
	log.Println("NotificationService gRPC running on port 50054")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
