package auth

import (
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
