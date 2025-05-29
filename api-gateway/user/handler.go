package user

import (
	"context"
	"net/http"

	pb "github.com/gemdivk/LUMERA-SPA/user-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func handleGrpcError(c *gin.Context, err error) {
	st, ok := status.FromError(err)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	switch st.Code() {
	case codes.InvalidArgument:
		c.JSON(http.StatusBadRequest, gin.H{"error": st.Message()})
	case codes.NotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": st.Message()})
	case codes.AlreadyExists:
		c.JSON(http.StatusConflict, gin.H{"error": st.Message()})
	case codes.PermissionDenied:
		c.JSON(http.StatusForbidden, gin.H{"error": st.Message()})
	case codes.Unauthenticated:
		c.JSON(http.StatusUnauthorized, gin.H{"error": st.Message()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
	}
}

func Register(c *gin.Context) {
	var req pb.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	resp, err := UserClient.Register(c, &req)
	if err != nil {
		handleGrpcError(c, err)
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func Login(c *gin.Context) {
	var req pb.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login data"})
		return
	}
	resp, err := UserClient.Login(c, &req)
	if err != nil {
		handleGrpcError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func GetMe(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}
	md := metadata.New(map[string]string{"authorization": auth})
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	resp, err := UserClient.GetMe(ctx, &emptypb.Empty{})
	if err != nil {
		handleGrpcError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func GetProfile(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}
	md := metadata.New(map[string]string{"authorization": auth})
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	userID := c.Param("id")
	resp, err := UserClient.GetProfile(ctx, &pb.GetProfileRequest{UserId: userID})
	if err != nil {
		handleGrpcError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func UpdateProfile(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}
	md := metadata.New(map[string]string{"authorization": auth})
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	var req pb.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid update data"})
		return
	}
	resp, err := UserClient.UpdateProfile(ctx, &req)
	if err != nil {
		handleGrpcError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func AssignRole(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}
	md := metadata.New(map[string]string{"authorization": auth})
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	var req pb.AssignRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}
	resp, err := UserClient.AssignRole(ctx, &req)
	if err != nil {
		handleGrpcError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func RemoveRole(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}
	md := metadata.New(map[string]string{"authorization": auth})
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	var req pb.RemoveRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}
	resp, err := UserClient.RemoveRole(ctx, &req)
	if err != nil {
		handleGrpcError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ListRoles(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}
	md := metadata.New(map[string]string{"authorization": auth})
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	userID := c.Param("id")
	resp, err := UserClient.ListRoles(ctx, &pb.ListRolesRequest{UserId: userID})
	if err != nil {
		handleGrpcError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func GetAllUsers(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}
	md := metadata.New(map[string]string{"authorization": auth})
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	resp, err := UserClient.GetAllUsers(ctx, &emptypb.Empty{})
	if err != nil {
		handleGrpcError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func SearchUsers(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}
	md := metadata.New(map[string]string{"authorization": auth})
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	query := c.Query("q")
	resp, err := UserClient.SearchUsers(ctx, &pb.SearchUserRequest{Query: query})
	if err != nil {
		handleGrpcError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteUser(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}
	md := metadata.New(map[string]string{"authorization": auth})
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	userID := c.Param("id")
	resp, err := UserClient.DeleteUser(ctx, &pb.DeleteUserRequest{UserId: userID})
	if err != nil {
		handleGrpcError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.String(http.StatusBadRequest, "Missing token")
		return
	}
	_, err := UserClient.VerifyEmail(context.Background(), &pb.VerifyEmailRequest{Token: token})
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid or expired token")
		return
	}
	c.String(http.StatusOK, "Your email has been successfully verified!")
}
