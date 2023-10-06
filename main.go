package main

import (
	"log"
	"net/http"
	"russianwords/routes"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	mux := http.NewServeMux()
	routes.SetupRoutes(mux)

	// Configuring CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allows requests from React dev server assuming it runs on port 3000
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	// Applying the CORS middleware to our routes (mux)
	handler := c.Handler(mux)

	// Running the server
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
