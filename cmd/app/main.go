package main

import (
	"chess/internal/database"
	"chess/internal/user/user_handlers"
	"chess/internal/web/router"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	// DB
	db := database.Connect()

	// HTTP server and router
	newRouter := router.Router()
	user_handlers.InitRouting(newRouter, db)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allows all origins, you can restrict this to specific domains
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	routerWithCors := c.Handler(newRouter)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", routerWithCors)

	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
