package tests

import (
	"context"
	"testing"

	pb "github.com/gemdivk/LUMERA-SPA/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var adminClient pb.UserServiceClient
var adminToken string
var targetUserID = "f211972b-6a22-49ef-8c27-cdf3d3a0c0e1"

func setupAdminClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	adminClient = pb.NewUserServiceClient(conn)
}

func adminCtx() context.Context {
	ctx := context.Background()
	md := metadata.New(map[string]string{"authorization": "Bearer " + adminToken})
	return metadata.NewOutgoingContext(ctx, md)
}

func TestAdminLogin(t *testing.T) {
	setupAdminClient(t)
	resp, err := adminClient.Login(context.Background(), &pb.LoginRequest{
		Email:    "dauka@example.com",
		Password: "newpassword",
	})
	if err != nil {
		t.Fatalf("Admin Login failed: %v", err)
	}
	adminToken = resp.Token
	t.Log("Admin login successful")
}

func TestAdminAssignRole(t *testing.T) {
	if targetUserID == "" {
		t.Skip("No targetUserID set for AssignRole test")
	}
	_, err := adminClient.AssignRole(adminCtx(), &pb.AssignRoleRequest{
		UserId:   targetUserID,
		RoleName: "specialist",
	})
	if err != nil {
		t.Fatalf("Admin AssignRole failed: %v", err)
	}
	t.Log("Admin assigned specialist role")
}

func TestAdminRemoveRole(t *testing.T) {
	if targetUserID == "" {
		t.Skip("No targetUserID set for RemoveRole test")
	}
	_, err := adminClient.RemoveRole(adminCtx(), &pb.RemoveRoleRequest{
		UserId:   targetUserID,
		RoleName: "specialist",
	})
	if err != nil {
		t.Fatalf("Admin RemoveRole failed: %v", err)
	}
	t.Log("Admin removed specialist role")
}

func TestAdminGetAllUsers(t *testing.T) {
	resp, err := adminClient.GetAllUsers(adminCtx(), &emptypb.Empty{})
	if err != nil {
		t.Fatalf("GetAllUsers failed: %v", err)
	}
	if len(resp.Users) == 0 {
		t.Fatalf("Expected users, got none")
	}
	t.Logf("Total users: %d", len(resp.Users))
}

func TestAdminSearchUsers(t *testing.T) {
	resp, err := adminClient.SearchUsers(adminCtx(), &pb.SearchUserRequest{Query: "testuser"})
	if err != nil {
		t.Fatalf("SearchUsers failed: %v", err)
	}
	if len(resp.Users) == 0 {
		t.Log("No users found for 'testuser'")
	} else {
		t.Logf("Found %d users matching 'testuser'", len(resp.Users))
	}
}

func TestAdminDeleteUser(t *testing.T) {
	if targetUserID == "" {
		t.Skip("No targetUserID set for DeleteUser test")
	}
	_, err := adminClient.DeleteUser(adminCtx(), &pb.DeleteUserRequest{UserId: targetUserID})
	if err != nil {
		t.Fatalf("Admin DeleteUser failed: %v", err)
	}
	t.Log("Admin deleted user")
}
