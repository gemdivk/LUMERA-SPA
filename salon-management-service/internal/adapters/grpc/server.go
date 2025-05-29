package grpc

import (
	"context"

	"github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/application"
	"github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/entity"
	pb "github.com/gemdivk/LUMERA-SPA/salon-management-service/proto"
)

type SalonServer struct {
	pb.UnimplementedSalonServiceServer
	uc application.SalonUsecase
}

func NewSalonServer(uc application.SalonUsecase) *SalonServer {
	return &SalonServer{uc: uc}
}

func (s *SalonServer) AddSalon(ctx context.Context, req *pb.AddSalonRequest) (*pb.SalonResponse, error) {
	result, err := s.uc.AddSalon(&entity.Salon{
		Name:     req.Name,
		Location: req.Location,
	})
	if err != nil {
		return nil, err
	}
	return &pb.SalonResponse{
		Id:       result.ID,
		Name:     result.Name,
		Location: result.Location,
	}, nil
}

func (s *SalonServer) AddProcedure(ctx context.Context, req *pb.AddProcedureRequest) (*pb.ProcedureResponse, error) {
	result, err := s.uc.AddProcedure(&entity.Procedure{
		SalonID:  req.SalonId,
		Name:     req.Name,
		Duration: req.Duration,
	})
	if err != nil {
		return nil, err
	}
	return &pb.ProcedureResponse{
		Id:       result.ID,
		SalonId:  result.SalonID,
		Name:     result.Name,
		Duration: result.Duration,
	}, nil
}

func (s *SalonServer) AddSpecialist(ctx context.Context, req *pb.AddSpecialistRequest) (*pb.SpecialistResponse, error) {
	result, err := s.uc.AddSpecialist(&entity.Specialist{
		SalonID: req.SalonId,
		Name:    req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &pb.SpecialistResponse{
		Id:      result.ID,
		SalonId: result.SalonID,
		Name:    result.Name,
	}, nil
}
