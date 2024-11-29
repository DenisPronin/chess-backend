package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
	"time"
)

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		authHeader := request.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(writer, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		secretKey := os.Getenv("JWT_SECRET_KEY")
		if secretKey == "" {
			http.Error(writer, "Secret key not set in environment variables", http.StatusInternalServerError)
			return
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(writer, "Invalid token", http.StatusUnauthorized)
			return
		}

		exp, ok := claims["exp"].(float64)
		if ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				http.Error(writer, "Token has expired", http.StatusUnauthorized)
				return
			}
		}

		ctx := context.WithValue(request.Context(), "userID", claims["userID"])

		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
