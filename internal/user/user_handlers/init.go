package user_handlers

import (
	"chess/internal/user"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func InitRouting(router *mux.Router, db *pgxpool.Pool) {
	userRepo := user.NewRepositoryUser(db)

	userHandler := &UserHandler{
		UserRepo: userRepo,
	}

	router.HandleFunc("/register", userHandler.Register).Methods(http.MethodPost)

	router.HandleFunc("/login", userHandler.Login).Methods(http.MethodPost)
}
