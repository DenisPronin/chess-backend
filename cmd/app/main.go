package main

import (
	"chess/internal/database"
	"chess/internal/user/user_handlers"
	"chess/internal/web/router"
	"log"
	"net/http"
)

func main() {
	// DB
	db := database.Connect()

	// HTTP server and router
	newRouter := router.Router()
	user_handlers.InitRouting(newRouter, db)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", newRouter)

	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
