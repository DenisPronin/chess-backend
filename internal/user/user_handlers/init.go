package user_handlers

import (
	"chess/internal/user"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func InitHandlers(router *mux.Router, db *pgxpool.Pool) {
	userRepo := user.NewRepositoryUser(db)

	registerHandler := &RegisterHandler{
		UserRepo: userRepo,
	}

	router.HandleFunc("/register", registerHandler.Register).Methods(http.MethodPost)

	router.HandleFunc("/login", Login).Methods(http.MethodPost)
}
