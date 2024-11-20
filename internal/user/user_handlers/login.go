package user_handlers

import (
	"chess/internal/user"
	"log"
	"net/http"
)

type LoginHandler struct {
	UserRepo *user.RepositoryUser
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *LoginHandler) Login(writer http.ResponseWriter, request *http.Request) {
	log.Println("Login", request)
}
