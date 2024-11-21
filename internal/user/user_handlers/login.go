package user_handlers

import (
	"context"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *UserHandler) Login(writer http.ResponseWriter, request *http.Request) {
	var req LoginRequest

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		http.Error(writer, "Invalid request format", http.StatusBadRequest)
		return
	}

	if req.Password == "" {
		http.Error(writer, "Missing required fields", http.StatusBadRequest)
		return
	}

	if req.Email == "" && req.Username == "" {
		http.Error(writer, "Missing required fields", http.StatusBadRequest)
		return
	}

	user, err := h.UserRepo.GetUserByUsernameOrEmail(context.Background(), req.Email, req.Username)
	if err != nil {
		http.Error(writer, "Invalid username/email or password", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		http.Error(writer, "Invalid username/email or password", http.StatusUnauthorized)
		return
	}

	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write([]byte("Login successful"))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
