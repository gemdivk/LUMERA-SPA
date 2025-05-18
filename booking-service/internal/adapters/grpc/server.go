package grpc

import (
	"context"

	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/usecase"
	pb "github.com/gemdivk/LUMERA-SPA/booking-service/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookingServer struct {
	pb.UnimplementedBookingServiceServer
	Usecase *usecase.BookingInteractor
}

func NewBookingServer(uc *usecase.BookingInteractor) *BookingServer {
	return &BookingServer{Usecase: uc}
}

// === CLIENT ===

func (s *BookingServer) ListAvailableSlots(ctx context.Context, req *pb.ListAvailableSlotsRequest) (*pb.ListAvailableSlotsResponse, error) {
	slots, err := s.Usecase.ListAvailableSlots(req.ProcedureId, req.Date)
	if err != nil {
		return nil, err
	}
	var pbSlots []*pb.TimeSlot
	for _, slot := range slots {
		pbSlots = append(pbSlots, &pb.TimeSlot{
			SpecialistId:   slot.SpecialistID,
			SpecialistName: slot.SpecialistName,
			StartTime:      timestamppb.New(slot.StartTime),
			EndTime:        timestamppb.New(slot.EndTime),
		})
	}
	return &pb.ListAvailableSlotsResponse{Slots: pbSlots}, nil
}

func (s *BookingServer) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.BookingResponse, error) {
	start := req.StartTime.AsTime()
	booking, err := s.Usecase.CreateBooking(req.ClientId, req.SpecialistId, req.ProcedureId, start)
	if err != nil {
		return nil, err
	}
	return &pb.BookingResponse{
		BookingId: booking.ID,
		Status:    booking.Status,
	}, nil
}

func (s *BookingServer) CancelBooking(ctx context.Context, req *pb.CancelBookingRequest) (*pb.CancelBookingResponse, error) {
	err := s.Usecase.CancelBooking(req.BookingId)
	return &pb.CancelBookingResponse{Success: err == nil}, err
}

func (s *BookingServer) ListClientBookings(ctx context.Context, req *pb.ClientBookingsRequest) (*pb.ClientBookingsResponse, error) {
	bookings, err := s.Usecase.GetClientBookings(req.ClientId)
	if err != nil {
		return nil, err
	}
	var pbBookings []*pb.Booking
	for _, b := range bookings {
		pbBookings = append(pbBookings, &pb.Booking{
			Id:             b.ID,
			ProcedureName:  "TODO", // Optional: join procedure name
			SpecialistName: "TODO", // Optional: join specialist name
			StartTime:      timestamppb.New(b.StartTime),
			EndTime:        timestamppb.New(b.EndTime),
			Status:         b.Status,
		})
	}
	return &pb.ClientBookingsResponse{Bookings: pbBookings}, nil
}

// === ADMIN ===

func (s *BookingServer) AdminCreateProcedure(ctx context.Context, req *pb.CreateProcedureRequest) (*pb.ProcedureResponse, error) {
	p, err := s.Usecase.CreateProcedure(req.Name, req.DurationMinutes)
	if err != nil {
		return nil, err
	}
	return &pb.ProcedureResponse{
		Id:              p.ID,
		Name:            p.Name,
		DurationMinutes: p.DurationMinutes,
	}, nil
}

func (s *BookingServer) AdminAssignProcedure(ctx context.Context, req *pb.AssignProcedureRequest) (*pb.AssignProcedureResponse, error) {
	err := s.Usecase.AssignProcedure(req.SpecialistId, req.ProcedureId)
	return &pb.AssignProcedureResponse{Success: err == nil}, err
}

func (s *BookingServer) AdminCreateScheduleTemplate(ctx context.Context, req *pb.CreateScheduleTemplateRequest) (*pb.ScheduleTemplateResponse, error) {
	err := s.Usecase.CreateScheduleTemplate(req.SpecialistId, int(req.Weekday), req.StartTime, req.EndTime, req.BreakMinutes)
	return &pb.ScheduleTemplateResponse{Success: err == nil}, err
}

func (s *BookingServer) AdminOverrideDaySchedule(ctx context.Context, req *pb.OverrideDayScheduleRequest) (*pb.OverrideDayScheduleResponse, error) {
	err := s.Usecase.OverrideDaySchedule(req.SpecialistId, req.Date, req.StartTime, req.EndTime, req.IsCancelled)
	return &pb.OverrideDayScheduleResponse{Success: err == nil}, err
}
