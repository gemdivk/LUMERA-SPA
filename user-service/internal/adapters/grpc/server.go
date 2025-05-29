package grpc

import (
	"context"
	"log"
	"time"

	"github.com/gemdivk/LUMERA-SPA/user-service/internal/domain/application"
	"github.com/gemdivk/LUMERA-SPA/user-service/internal/usecase"
	pb "github.com/gemdivk/LUMERA-SPA/user-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	Usecase application.UserUsecase
}

func NewUserServer(uc *usecase.UserInteractor) *UserServer {
	return &UserServer{Usecase: uc}
}

func (s *UserServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "name, email and password are required")
	}

	user, token, err := s.Usecase.Register(req.Name, req.Email, req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.AuthResponse{
		Token: token,
		Profile: &pb.UserProfile{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
		},
	}, nil
}

func (s *UserServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "email and password are required")
	}

	token, err := s.Usecase.Login(req.Email, req.Password)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	return &pb.AuthResponse{
		Token: token,
	}, nil
}

func (s *UserServer) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.UserProfile, error) {
	user, err := s.Usecase.GetProfile(req.UserId)
	if err != nil {
		return nil, err
	}
	roles, _ := s.Usecase.ListRoles(req.UserId)
	return &pb.UserProfile{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		Roles:     roles,
	}, nil
}

func (s *UserServer) GetMe(ctx context.Context, _ *emptypb.Empty) (*pb.UserProfile, error) {
	claims, err := extractJWTClaims(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	user, err := s.Usecase.GetProfile(claims.UserID)
	if err != nil {
		return nil, err
	}
	roles, _ := s.Usecase.ListRoles(claims.UserID)
	return &pb.UserProfile{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		Roles:     roles,
	}, nil
}

func (s *UserServer) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UserProfile, error) {
	claims, err := extractJWTClaims(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	user, err := s.Usecase.UpdateProfile(claims.UserID, req.Name, req.Password)
	if err != nil {
		return nil, err
	}
	roles, _ := s.Usecase.ListRoles(claims.UserID)
	return &pb.UserProfile{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		Roles:     roles,
	}, nil
}

func (s *UserServer) AssignRole(ctx context.Context, req *pb.AssignRoleRequest) (*pb.AssignRoleResponse, error) {
	claims, err := extractJWTClaims(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if !isAdmin(claims) {
		return nil, status.Error(codes.PermissionDenied, "only admins can assign roles")
	}
	err = s.Usecase.AssignRole(req.UserId, req.RoleName)
	return &pb.AssignRoleResponse{Success: err == nil}, err
}

func (s *UserServer) ListRoles(ctx context.Context, req *pb.ListRolesRequest) (*pb.ListRolesResponse, error) {
	roles, err := s.Usecase.ListRoles(req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.ListRolesResponse{Roles: roles}, nil
}

func (s *UserServer) GetAllUsers(ctx context.Context, _ *emptypb.Empty) (*pb.UserList, error) {
	claims, err := extractJWTClaims(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if !isAdmin(claims) {
		return nil, status.Error(codes.PermissionDenied, "admin only")
	}

	users, err := s.Usecase.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var pbUsers []*pb.UserProfile
	for _, user := range users {
		roles, _ := s.Usecase.ListRoles(user.ID)
		pbUsers = append(pbUsers, &pb.UserProfile{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			Roles:     roles,
		})
	}
	return &pb.UserList{Users: pbUsers}, nil
}

func (s *UserServer) SearchUsers(ctx context.Context, req *pb.SearchUserRequest) (*pb.UserList, error) {
	claims, err := extractJWTClaims(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if !isAdmin(claims) {
		return nil, status.Error(codes.PermissionDenied, "admin only")
	}

	users, err := s.Usecase.SearchUsers(req.Query)
	if err != nil {
		return nil, err
	}

	var pbUsers []*pb.UserProfile
	for _, user := range users {
		roles, _ := s.Usecase.ListRoles(user.ID)
		pbUsers = append(pbUsers, &pb.UserProfile{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			Roles:     roles,
		})
	}
	return &pb.UserList{Users: pbUsers}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	claims, err := extractJWTClaims(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if !isAdmin(claims) {
		return nil, status.Error(codes.PermissionDenied, "admin only")
	}
	if claims.UserID == req.UserId {
		return nil, status.Error(codes.FailedPrecondition, "you cannot delete your own account")
	}

	err = s.Usecase.DeleteUser(req.UserId)
	return &pb.DeleteUserResponse{Success: err == nil}, err
}
func (s *UserServer) RemoveRole(ctx context.Context, req *pb.RemoveRoleRequest) (*pb.RemoveRoleResponse, error) {
	log.Printf("[GRPC] RemoveRole called with user_id=%s, role_name=%s", req.GetUserId(), req.GetRoleName())
	claims, err := extractJWTClaims(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if !isAdmin(claims) {
		return nil, status.Error(codes.PermissionDenied, "only admins can remove roles")
	}
	err = s.Usecase.RemoveRole(req.UserId, req.RoleName)
	return &pb.RemoveRoleResponse{Success: err == nil}, err
}
func (s *UserServer) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	err := s.Usecase.MarkEmailVerified(req.Token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid token")
	}
	return &pb.VerifyEmailResponse{Success: true}, nil
}
func (s *UserServer) Logout(ctx context.Context, _ *emptypb.Empty) (*pb.LogoutResponse, error) {
	_, err := extractJWTClaims(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	return &pb.LogoutResponse{Success: true}, nil
}
