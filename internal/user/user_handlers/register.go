package user_handlers

import (
	"chess/internal/user"
	"context"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type RegisterHandler struct {
	UserRepo *user.RepositoryUser
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *RegisterHandler) Register(writer http.ResponseWriter, request *http.Request) {
	var req RegisterRequest

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		http.Error(writer, "Invalid request format", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Username == "" || req.Password == "" {
		http.Error(writer, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Проверить наличие юзера

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(writer, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	userModel := user.User{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	err = h.UserRepo.SaveUser(context.Background(), userModel)
	if err != nil {
		http.Error(writer, "Failed to save user", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("User registered successfully"))
}
