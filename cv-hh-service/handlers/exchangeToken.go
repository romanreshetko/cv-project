package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

const clientId = "G69FEKNJM1JRUBA0G48L8TAFRJBEO4E0DJIR3B4Q87FN87GU1NGA67BO8SJJE8D2"
const clientSecret = "IDO93HOAKJ8KL48MS229AJR0LKSUEMGC02K5312I755U5L0H1QAB7NR8PEHSFN70"
const redirectUri = "https://husky-notable-jackal.ngrok-free.app/hh-upload"

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type ExchangeRequest struct {
	Code string `json:"code"`
}

func ExchangeToken(w http.ResponseWriter, r *http.Request) {
	var reqBody ExchangeRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		writeJSONError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("code", reqBody.Code)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", redirectUri)

	resp, err := http.Post(
		"https://api.hh.ru/token",
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(data.Encode()),
	)

	if err != nil {
		writeJSONError(w, "Error while requesting token", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		writeJSONError(w, "Error while requesting token", resp.StatusCode)
		return
	}

	var tokenResponse TokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		writeJSONError(w, "Failed to parse token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenResponse)
}

func writeJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}
