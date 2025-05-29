package grpc

import (
	"context"
	"github.com/gemdivk/LUMERA-SPA/payment-service/internal/usecase"
	pb "github.com/gemdivk/LUMERA-SPA/payment-service/proto"
)

type Server struct {
	pb.UnimplementedPaymentServiceServer
	UC *usecase.PaymentUsecase
}

func (s *Server) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.PaymentResponse, error) {
	p, err := s.UC.Create(ctx, req.UserId, req.Amount, req.Currency, req.PaymentMethod, req.SalonId, req.ServiceId)
	if err != nil {
		return nil, err
	}
	return &pb.PaymentResponse{Id: p.ID, Status: p.Status}, nil
}

func (s *Server) ListPayments(ctx context.Context, req *pb.UserRequest) (*pb.PaymentListResponse, error) {
	payments, err := s.UC.List(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	var res []*pb.Payment
	for _, p := range payments {
		res = append(res, &pb.Payment{
			Id: p.ID, UserId: p.UserID, Amount: p.Amount, Currency: p.Currency,
			PaymentMethod: p.PaymentMethod, Status: p.Status, CreatedAt: p.CreatedAt,
		})
	}
	return &pb.PaymentListResponse{Payments: res}, nil
}
