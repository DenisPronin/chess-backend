package user_handlers

import (
	"chess/internal/json_errors"
	"chess/internal/user"
	"context"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *UserHandler) Register(writer http.ResponseWriter, request *http.Request) {
	var req RegisterRequest

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		json_errors.CatchError(writer, http.StatusBadRequest, "Invalid request format")
		return
	}

	if req.Email == "" || req.Username == "" || req.Password == "" {
		json_errors.CatchError(writer, http.StatusBadRequest, "Missing required fields")
		return
	}

	// Проверить наличие юзера

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		json_errors.CatchError(writer, http.StatusInternalServerError, "Failed to hash password")
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
		json_errors.CatchError(writer, http.StatusInternalServerError, "Failed to save user")
		return
	}

	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write([]byte("User registered successfully"))
	if err != nil {
		json_errors.CatchError(writer, http.StatusInternalServerError, "Error writing response")
	}
}
