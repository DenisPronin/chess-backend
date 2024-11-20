package router

import (
	"chess/internal/health"
	"chess/internal/user"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", user.HandlerLogin).Methods(http.MethodPost)

	router.HandleFunc("/health", health.HandlerHealth).Methods(http.MethodGet)

	return router
}
