package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	grpcadapter "github.com/gemdivk/LUMERA-SPA/booking-service/internal/adapters/grpc"
	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/infrastructure/postgres"
	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/usecase"
	pb "github.com/gemdivk/LUMERA-SPA/booking-service/proto"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Failed to load .env")
	}

	db, err := sql.Open("postgres", os.Getenv("BOOKING_DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bookingRepo := postgres.NewBookingRepo(db)
	procedureRepo := postgres.NewProcedureRepo(db)
	scheduleRepo := postgres.NewScheduleRepo(db)

	interactor := usecase.NewBookingInteractor(bookingRepo, procedureRepo, scheduleRepo)

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBookingServiceServer(grpcServer, grpcadapter.NewBookingServer(interactor))

	log.Println("BookingService running on port 50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
