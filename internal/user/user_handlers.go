package user

import (
	"log"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandlerLogin(writer http.ResponseWriter, request *http.Request) {
	log.Println("Login", request)
}
