package router

import (
	"chess/internal/health"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", health.HandlerHealth).Methods(http.MethodGet)

	return router
}
