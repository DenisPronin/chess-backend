package user_handlers

import (
	"chess/internal/auth"
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

	jwtToken, err := auth.GenerateJWT(user.ID, user.Username)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	response := map[string]string{"token": jwtToken}

	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
