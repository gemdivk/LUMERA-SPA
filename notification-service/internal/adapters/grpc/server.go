package grpc

import (
	"context"
	"time"

	"github.com/gemdivk/LUMERA-SPA/notification-service/internal/domain/application"
	pb "github.com/gemdivk/LUMERA-SPA/notification-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NotificationServer struct {
	pb.UnimplementedNotificationServiceServer
	Usecase application.EmailUsecase
}

func NewNotificationServer(uc application.EmailUsecase) *NotificationServer {
	return &NotificationServer{Usecase: uc}
}

func (s *NotificationServer) GetEmailLogs(ctx context.Context, _ *emptypb.Empty) (*pb.EmailLogList, error) {
	claims, err := extractJWTClaims(ctx)
	if err != nil || !isAdmin(claims) {
		return nil, status.Error(codes.PermissionDenied, "admin only")
	}
	logs, err := s.Usecase.GetLogs()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var pbLogs []*pb.EmailLog
	for _, l := range logs {
		pbLogs = append(pbLogs, &pb.EmailLog{
			Id:      int32(l.ID),
			Email:   l.Email,
			Subject: l.Subject,
			SentAt:  l.SentAt.Format(time.RFC3339),
		})
	}
	return &pb.EmailLogList{Logs: pbLogs}, nil
}
