package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(userID, email string, roles []string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"roles":   roles,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
func ParseToken(tokenStr string) (*Claims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return &Claims{
		UserID: claims["user_id"].(string),
		Email:  claims["email"].(string),
		Roles:  toStringSlice(claims["roles"]),
	}, nil
}

type Claims struct {
	UserID string
	Email  string
	Roles  []string
}

func toStringSlice(v interface{}) []string {
	raw, ok := v.([]interface{})
	if !ok {
		return nil
	}
	var result []string
	for _, val := range raw {
		if str, ok := val.(string); ok {
			result = append(result, str)
		}
	}
	return result
}
