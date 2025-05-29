package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID string
	Email  string
	Roles  []string
}

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
	parsedClaims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, parsedClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		if validationErr, ok := err.(*jwt.ValidationError); ok {
			if validationErr.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("token expired")
			}
		}
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	userID, ok1 := parsedClaims["user_id"].(string)
	email, ok2 := parsedClaims["email"].(string)
	rolesRaw := parsedClaims["roles"]

	if !ok1 || !ok2 {
		return nil, fmt.Errorf("invalid token claims")
	}

	return &Claims{
		UserID: userID,
		Email:  email,
		Roles:  toStringSlice(rolesRaw),
	}, nil
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
