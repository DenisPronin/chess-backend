package router

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

	return api
}
