package grpc

import (
	"context"
	"github.com/gemdivk/LUMERA-SPA/review-service/internal/domain"
	"github.com/gemdivk/LUMERA-SPA/review-service/internal/domain/application"
	pb "github.com/gemdivk/LUMERA-SPA/review-service/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ReviewServer struct {
	pb.UnimplementedReviewServiceServer
	Usecase application.ReviewUsecase
}

func NewReviewServer(uc application.ReviewUsecase) *ReviewServer {
	return &ReviewServer{Usecase: uc}
}

func (s *ReviewServer) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.ReviewResponse, error) {
	review, err := s.Usecase.CreateReview(&domain.Review{
		SalonID: req.SalonId,
		UserID:  req.UserId,
		Content: req.Content,
		Rating:  req.Rating,
	})
	if err != nil {
		return nil, err
	}
	return toProto(review), nil
}

func (s *ReviewServer) GetReview(ctx context.Context, req *pb.GetReviewRequest) (*pb.ReviewResponse, error) {
	review, err := s.Usecase.GetReview(req.Id)
	if err != nil {
		return nil, err
	}
	return toProto(review), nil
}

func (s *ReviewServer) UpdateReview(ctx context.Context, req *pb.UpdateReviewRequest) (*pb.ReviewResponse, error) {
	review, err := s.Usecase.UpdateReview(&domain.Review{
		ID:      req.Id,
		Content: req.Content,
		Rating:  req.Rating,
	})
	if err != nil {
		return nil, err
	}
	return toProto(review), nil
}

func (s *ReviewServer) DeleteReview(ctx context.Context, req *pb.DeleteReviewRequest) (*pb.DeleteReviewResponse, error) {
	err := s.Usecase.DeleteReview(req.Id)
	return &pb.DeleteReviewResponse{Success: err == nil}, err
}

func (s *ReviewServer) ListReviews(ctx context.Context, req *pb.ListReviewsRequest) (*pb.ListReviewsResponse, error) {
	list, err := s.Usecase.ListBySalon(req.SalonId)
	if err != nil {
		return nil, err
	}
	var resp []*pb.ReviewResponse
	for _, r := range list {
		resp = append(resp, toProto(r))
	}
	return &pb.ListReviewsResponse{Reviews: resp}, nil
}

func toProto(r *domain.Review) *pb.ReviewResponse {
	return &pb.ReviewResponse{
		Id:        r.ID,
		SalonId:   r.SalonID,
		UserId:    r.UserID,
		Content:   r.Content,
		Rating:    r.Rating,
		CreatedAt: timestamppb.New(r.CreatedAt),
	}
}
