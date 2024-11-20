package user_handlers

import (
	"log"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Login(writer http.ResponseWriter, request *http.Request) {
	log.Println("Login", request)
}
