package grpc

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/metadata"
)

type Claims struct {
	UserID string
	Email  string
	Roles  []string
}

func extractJWTClaims(ctx context.Context) (*Claims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("metadata not found")
	}
	authHeaders := md.Get("authorization")
	if len(authHeaders) == 0 {
		return nil, fmt.Errorf("authorization header not found")
	}
	tokenStr := strings.TrimPrefix(authHeaders[0], "Bearer ")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid or expired token")
	}
	claimsMap, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	var roles []string
	if rawRoles, ok := claimsMap["roles"].([]interface{}); ok {
		for _, r := range rawRoles {
			if str, ok := r.(string); ok {
				roles = append(roles, str)
			}
		}
	}
	return &Claims{
		UserID: fmt.Sprintf("%v", claimsMap["user_id"]),
		Email:  fmt.Sprintf("%v", claimsMap["email"]),
		Roles:  roles,
	}, nil
}

func isAdmin(c *Claims) bool {
	for _, r := range c.Roles {
		if r == "admin" {
			return true
		}
	}
	return false
}
