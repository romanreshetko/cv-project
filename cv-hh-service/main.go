package main

import (
	"cv-hh-service/handlers"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/exchange-token", handlers.ExchangeToken)
	mux.HandleFunc("/update", handlers.UpdateResumeHandler)
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handler := corsHandler.Handler(mux)
	log.Fatal(http.ListenAndServe(":8082", handler))
}
