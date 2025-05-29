package integration

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net"
	"os"
	"testing"
	"time"

	grpcAdapter "github.com/gemdivk/LUMERA-SPA/user-service/internal/adapters/grpc"
	"github.com/gemdivk/LUMERA-SPA/user-service/internal/infrastructure/cache"
	"github.com/gemdivk/LUMERA-SPA/user-service/internal/infrastructure/postgres"
	"github.com/gemdivk/LUMERA-SPA/user-service/internal/usecase"
	pb "github.com/gemdivk/LUMERA-SPA/user-service/proto"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	client     pb.UserServiceClient
	adminToken string
	userToken  string
	userID     string
	nc         *nats.Conn
)

func TestMain(m *testing.M) {
	_ = godotenv.Load("../../.env")

	connStr := "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	nc, err = nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		log.Fatal(err)
	}

	repo := postgres.NewUserRepo(db)
	userCache := cache.NewUserCache()
	users, _ := repo.GetAll()
	userCache.LoadInitial(users)

	uc := usecase.NewUserInteractorWithCache(repo, nc, userCache)
	srv := grpcAdapter.NewUserServer(uc)

	go func() {
		lis, err := net.Listen("tcp", ":50570")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		server := grpc.NewServer()
		pb.RegisterUserServiceServer(server, srv)
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	time.Sleep(3 * time.Second)

	conn, err := grpc.Dial("localhost:50570", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial gRPC server: %v", err)
	}

	client = pb.NewUserServiceClient(conn)
	os.Exit(m.Run())
}

func TestUserFlow(t *testing.T) {
	ctx := context.Background()

	email := "intg@example.com"
	password := "pass123"
	name := "Intg User"

	// Register
	registerResp, err := client.Register(ctx, &pb.RegisterRequest{
		Name:     name,
		Email:    email,
		Password: password,
	})
	require.NoError(t, err)
	require.NotNil(t, registerResp)
	require.NotNil(t, registerResp.Profile)
	userID = registerResp.Profile.Id

	// Intercept NATS email
	sub, _ := nc.SubscribeSync("notifications.email.verification")
	msg, err := sub.NextMsg(2 * time.Second)
	require.NoError(t, err)

	var payload map[string]string
	err = json.Unmarshal(msg.Data, &payload)
	require.NoError(t, err)
	require.NotEmpty(t, payload["token"])

	// Verify email
	verifyResp, err := client.VerifyEmail(ctx, &pb.VerifyEmailRequest{Token: payload["token"]})
	require.NoError(t, err)
	require.True(t, verifyResp.Success)

	// Login
	loginResp, err := client.Login(ctx, &pb.LoginRequest{
		Email:    email,
		Password: password,
	})
	require.NoError(t, err)
	require.NotEmpty(t, loginResp.Token)
	userToken = loginResp.Token

	// GetProfile
	profile, err := client.GetProfile(ctx, &pb.GetProfileRequest{UserId: userID})
	require.NoError(t, err)
	assert.Equal(t, name, profile.Name)

	// UpdateProfile
	ctxWithToken := injectToken(ctx, userToken)
	updated, err := client.UpdateProfile(ctxWithToken, &pb.UpdateProfileRequest{
		Name:     "Updated",
		Password: "newpass",
	})
	require.NoError(t, err)
	assert.Equal(t, "Updated", updated.Name)

	// GetMe
	me, err := client.GetMe(ctxWithToken, &emptypb.Empty{})
	require.NoError(t, err)
	assert.Equal(t, "Updated", me.Name)

	// Admin actions
	adminToken = userToken
	adminCtx := injectToken(ctx, adminToken)

	// AssignRole
	assignResp, err := client.AssignRole(adminCtx, &pb.AssignRoleRequest{
		UserId:   userID,
		RoleName: "tester",
	})
	require.NoError(t, err)
	assert.True(t, assignResp.Success)

	// ListRoles
	roleList, err := client.ListRoles(adminCtx, &pb.ListRolesRequest{UserId: userID})
	require.NoError(t, err)
	assert.Contains(t, roleList.Roles, "tester")

	// SearchUsers
	search, err := client.SearchUsers(adminCtx, &pb.SearchUserRequest{Query: "intg"})
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(search.Users), 1)

	// GetAllUsers
	all, err := client.GetAllUsers(adminCtx, &emptypb.Empty{})
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(all.Users), 1)

	// RemoveRole
	removeResp, err := client.RemoveRole(adminCtx, &pb.RemoveRoleRequest{
		UserId:   userID,
		RoleName: "tester",
	})
	require.NoError(t, err)
	assert.True(t, removeResp.Success)

	// DeleteUser
	delResp, err := client.DeleteUser(adminCtx, &pb.DeleteUserRequest{UserId: userID})
	require.NoError(t, err)
	assert.True(t, delResp.Success)
}

func injectToken(ctx context.Context, token string) context.Context {
	md := metadata.New(map[string]string{"authorization": "Bearer " + token})
	return metadata.NewOutgoingContext(ctx, md)
}
