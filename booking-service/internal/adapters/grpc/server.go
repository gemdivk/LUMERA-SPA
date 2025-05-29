package grpc

import (
	"context"

	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain"
	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain/application"
	pb "github.com/gemdivk/LUMERA-SPA/booking-service/proto"
)

type BookingServer struct {
	pb.UnimplementedBookingServiceServer
	uc application.BookingUsecase
}

func NewBookingServer(uc application.BookingUsecase) *BookingServer {
	return &BookingServer{uc: uc}
}

func (s *BookingServer) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.BookingResponse, error) {
	b, err := s.uc.Create(&domain.Booking{
		ClientID:     req.ClientId,
		SalonID:      req.SalonId,
		ProcedureID:  req.ProcedureId,
		SpecialistID: req.SpecialistId,
		Date:         req.Date,
		StartTime:    req.StartTime,
	})
	if err != nil {
		return nil, err
	}
	return toProto(b), nil
}

func (s *BookingServer) CancelBooking(ctx context.Context, req *pb.CancelBookingRequest) (*pb.SuccessResponse, error) {
	err := s.uc.Cancel(req.BookingId)
	return &pb.SuccessResponse{Success: err == nil}, err
}

func (s *BookingServer) RescheduleBooking(ctx context.Context, req *pb.RescheduleBookingRequest) (*pb.BookingResponse, error) {
	b, err := s.uc.Reschedule(req.BookingId, req.NewDate, req.NewStartTime)
	if err != nil {
		return nil, err
	}
	return toProto(b), nil
}

func (s *BookingServer) ListBookings(ctx context.Context, req *pb.ListBookingsRequest) (*pb.ListBookingsResponse, error) {
	bookings, err := s.uc.ListByClient(req.ClientId)
	if err != nil {
		return nil, err
	}
	var pbBookings []*pb.BookingResponse
	for _, b := range bookings {
		pbBookings = append(pbBookings, toProto(b))
	}
	return &pb.ListBookingsResponse{Bookings: pbBookings}, nil
}

func toProto(b *domain.Booking) *pb.BookingResponse {
	return &pb.BookingResponse{
		Id:           b.ID,
		ClientId:     b.ClientID,
		SalonId:      b.SalonID,
		ProcedureId:  b.ProcedureID,
		SpecialistId: b.SpecialistID,
		Date:         b.Date,
		StartTime:    b.StartTime,
		Status:       b.Status,
	}
}
