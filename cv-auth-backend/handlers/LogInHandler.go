package handlers

import (
	"cv-auth-backend/database"
	"cv-auth-backend/jwt"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type LogInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogInResponse struct {
	Token string `json:"token"`
}

func LogInHandler(w http.ResponseWriter, r *http.Request) {
	var req LogInRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	passwordHash, err := database.GetUserByUsername(req.Username)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(*passwordHash), []byte(req.Password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := jwt.GenerateJWT(req.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LogInResponse{Token: token})
}
