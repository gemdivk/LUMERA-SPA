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

func (s *SalonServer) GetAllSpecialists(ctx context.Context, _ *pb.Empty) (*pb.SpecialistListResponse, error) {
	specialists, err := s.uc.GetAllSpecialists()
	if err != nil {
		return nil, err
	}

	var res []*pb.SpecialistResponse
	for _, sp := range specialists {
		res = append(res, &pb.SpecialistResponse{
			Id:      sp.ID,
			SalonId: sp.SalonID,
			Name:    sp.Name,
		})
	}
	return &pb.SpecialistListResponse{Specialists: res}, nil
}

func (s *SalonServer) GetAllProcedures(ctx context.Context, _ *pb.Empty) (*pb.ProcedureListResponse, error) {
	procedures, err := s.uc.GetAllProcedures()
	if err != nil {
		return nil, err
	}

	var res []*pb.ProcedureResponse
	for _, pr := range procedures {
		res = append(res, &pb.ProcedureResponse{
			Id:       pr.ID,
			SalonId:  pr.SalonID,
			Name:     pr.Name,
			Duration: pr.Duration,
		})
	}
	return &pb.ProcedureListResponse{Procedures: res}, nil
}

func (s *SalonServer) AssignProcedureToSpecialist(ctx context.Context, req *pb.AssignProcedureRequest) (*pb.AssignResponse, error) {
	err := s.uc.AssignProcedureToSpecialist(req.SpecialistId, req.ProcedureId)
	if err != nil {
		return nil, err
	}
	return &pb.AssignResponse{Success: true}, nil
}
