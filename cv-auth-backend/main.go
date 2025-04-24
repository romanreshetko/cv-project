package main

import (
	"cv-auth-backend/database"
	"cv-auth-backend/handlers"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	err := database.ConnectDB()
	if err != nil {
		return
	}
	defer database.DB.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/signup", handlers.SignUpHandler)
	mux.HandleFunc("/login", handlers.LogInHandler)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handler := corsHandler.Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
