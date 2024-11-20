package main

import (
	"chess/internal/db"
	"chess/internal/web/router"
	"log"
	"net/http"
)

func main() {
	// DB
	db.Connect()

	// HTTP server and router
	r := router.Router()

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
