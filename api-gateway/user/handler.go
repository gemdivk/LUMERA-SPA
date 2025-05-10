package user

import (
	"net/http"

	pb "github.com/gemdivk/LUMERA-SPA/user-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Register(c *gin.Context) {
	var req pb.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	resp, err := UserClient.Register(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func Login(c *gin.Context) {
	var req pb.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login body"})
		return
	}
	resp, err := UserClient.Login(c, &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func GetMe(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no token"})
		return
	}
	md := metadata.New(map[string]string{"authorization": auth})
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)
	resp, err := UserClient.GetMe(ctx, &emptypb.Empty{})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func GetProfile(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no token"})
		return
	}
	md := metadata.New(map[string]string{"authorization": auth})
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	userID := c.Param("id")
	resp, err := UserClient.GetProfile(ctx, &pb.GetProfileRequest{UserId: userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func UpdateProfile(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no token"})
		return
	}
	md := metadata.New(map[string]string{"authorization": auth})
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	var req pb.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile data"})
		return
	}
	resp, err := UserClient.UpdateProfile(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func AssignRole(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no token"})
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
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func RemoveRole(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no token"})
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
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func ListRoles(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no token"})
		return
	}
	md := metadata.New(map[string]string{"authorization": auth})
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	userID := c.Param("id")
	resp, err := UserClient.ListRoles(ctx, &pb.ListRolesRequest{UserId: userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func GetAllUsers(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no token"})
		return
	}
	md := metadata.New(map[string]string{"authorization": auth})
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	resp, err := UserClient.GetAllUsers(ctx, &emptypb.Empty{})
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func SearchUsers(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no token"})
		return
	}
	md := metadata.New(map[string]string{"authorization": auth})
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	query := c.Query("q")
	resp, err := UserClient.SearchUsers(ctx, &pb.SearchUserRequest{Query: query})
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteUser(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no token"})
		return
	}
	md := metadata.New(map[string]string{"authorization": auth})
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	userID := c.Param("id")
	resp, err := UserClient.DeleteUser(ctx, &pb.DeleteUserRequest{UserId: userID})
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
