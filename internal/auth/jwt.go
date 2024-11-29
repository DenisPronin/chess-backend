package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

const tokenExpiry = 24 * time.Hour

func GenerateJWT(userID int, username string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("secret key not set in environment variables")
	}

	claims := jwt.MapClaims{
		"userID":   userID,
		"username": username,
		"exp":      time.Now().Add(tokenExpiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
