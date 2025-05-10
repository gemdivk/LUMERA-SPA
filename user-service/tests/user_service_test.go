package tests

import (
	"context"
	"testing"
	"time"

	pb "github.com/gemdivk/LUMERA-SPA/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var client pb.UserServiceClient
var token string
var testUserID string

func setupClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	client = pb.NewUserServiceClient(conn)
}

func authCtx() context.Context {
	ctx := context.Background()
	md := metadata.New(map[string]string{"authorization": "Bearer " + token})
	return metadata.NewOutgoingContext(ctx, md)
}

func TestRegister(t *testing.T) {
	setupClient(t)
	resp, err := client.Register(context.Background(), &pb.RegisterRequest{
		Name:     "TestUser",
		Email:    "testuser2025@example.com",
		Password: "testpass123",
	})
	if err != nil {
		t.Fatalf("Register failed: %v", err)
	}
	token = resp.Token
	testUserID = resp.Profile.Id
	t.Logf("Registered user with ID: %s", testUserID)
}

func TestLogin(t *testing.T) {
	resp, err := client.Login(context.Background(), &pb.LoginRequest{
		Email:    "testuser2025@example.com",
		Password: "testpass123",
	})
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}
	token = resp.Token
	t.Log("Login successful")
}

func TestGetMe(t *testing.T) {
	resp, err := client.GetMe(authCtx(), &emptypb.Empty{})
	if err != nil {
		t.Fatalf("GetMe failed: %v", err)
	}
	t.Logf("Me: %v", resp)
}

func TestUpdateProfile(t *testing.T) {
	resp, err := client.UpdateProfile(authCtx(), &pb.UpdateProfileRequest{
		Name:     "TestUserUpdated",
		Password: "newpass123",
	})
	if err != nil {
		t.Fatalf("UpdateProfile failed: %v", err)
	}
	if resp.Name != "TestUserUpdated" {
		t.Fatalf("Expected updated name, got: %s", resp.Name)
	}
	t.Log("UpdateProfile successful")
}

func TestAssignRole(t *testing.T) {
	_, err := client.AssignRole(authCtx(), &pb.AssignRoleRequest{
		UserId:   testUserID,
		RoleName: "specialist",
	})
	if err != nil {
		t.Logf("AssignRole (expected failure if not admin): %v", err)
	} else {
		t.Log("AssignRole success (admin confirmed)")
	}
}

func TestRemoveRole(t *testing.T) {
	_, err := client.RemoveRole(authCtx(), &pb.RemoveRoleRequest{
		UserId:   testUserID,
		RoleName: "specialist",
	})
	if err != nil {
		t.Logf("RemoveRole (expected failure if not admin): %v", err)
	} else {
		t.Log("RemoveRole success (admin confirmed)")
	}
}

func TestListRoles(t *testing.T) {
	resp, err := client.ListRoles(authCtx(), &pb.ListRolesRequest{UserId: testUserID})
	if err != nil {
		t.Fatalf("ListRoles failed: %v", err)
	}
	t.Logf("Roles: %v", resp.Roles)
}

func TestGetAllUsers(t *testing.T) {
	resp, err := client.GetAllUsers(authCtx(), &emptypb.Empty{})
	if err != nil {
		t.Logf("GetAllUsers (expected failure if not admin): %v", err)
		return
	}
	if len(resp.Users) == 0 {
		t.Fatalf("Expected users, got 0")
	}
	t.Logf("Total users: %d", len(resp.Users))
}

func TestDeleteUser(t *testing.T) {
	time.Sleep(2 * time.Second)
	_, err := client.DeleteUser(authCtx(), &pb.DeleteUserRequest{UserId: testUserID})
	if err != nil {
		t.Logf("DeleteUser (expected failure if not admin or deleting self): %v", err)
	} else {
		t.Log("DeleteUser success")
	}
}

func TestGetProfile(t *testing.T) {
	resp, err := client.GetProfile(authCtx(), &pb.GetProfileRequest{UserId: testUserID})
	if err != nil {
		t.Fatalf("GetProfile failed: %v", err)
	}
	if resp.Id != testUserID {
		t.Fatalf("Expected user_id %s, got %s", testUserID, resp.Id)
	}
	t.Logf("GetProfile success: %s", resp.Name)
}

func TestSearchUsers(t *testing.T) {
	resp, err := client.SearchUsers(authCtx(), &pb.SearchUserRequest{Query: "testuser"})
	if err != nil {
		t.Logf("SearchUsers (expected failure if not admin): %v", err)
		return
	}
	if len(resp.Users) == 0 {
		t.Logf("No users found for 'testuser'")
	} else {
		t.Logf("Found %d user(s)", len(resp.Users))
	}
}
