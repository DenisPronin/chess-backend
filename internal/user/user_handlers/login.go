package user_handlers

import (
	"chess/internal/auth"
	"chess/internal/json_errors"
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
		json_errors.CatchError(writer, http.StatusBadRequest, "Invalid request format")
		return
	}

	if req.Password == "" {
		json_errors.CatchError(writer, http.StatusBadRequest, "Missing required fields")
		return
	}

	if req.Email == "" && req.Username == "" {
		json_errors.CatchError(writer, http.StatusBadRequest, "Missing required fields")
		return
	}

	user, err := h.UserRepo.GetUserByUsernameOrEmail(context.Background(), req.Email, req.Username)
	if err != nil {
		json_errors.CatchError(writer, http.StatusUnauthorized, "Invalid username/email or password")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		json_errors.CatchError(writer, http.StatusUnauthorized, "Invalid username/email or password")
		return
	}

	jwtToken, err := auth.GenerateJWT(user.ID, user.Username)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		json_errors.CatchError(writer, http.StatusInternalServerError, "Internal server error")
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
