package handlers

import (
	. "cv-file-validate/models"
	. "cv-file-validate/validation"
	"encoding/json"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
)

const maxFileSize = 5 * 1024 * 1024

func ValidateFileHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxFileSize)
	err := r.ParseMultipartForm(maxFileSize)
	if err != nil {
		writeJSONError(w, "File is too large", http.StatusRequestEntityTooLarge)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		writeJSONError(w, "Invalid file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := fileHeader.Filename
	if !IsValidYAML(filename) {
		writeJSONError(w, "Only YAML files allowed", http.StatusUnprocessableEntity)
		return
	}

	fileContent, err := io.ReadAll(file)
	if err != nil {
		writeJSONError(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	var resume Resume
	err = yaml.Unmarshal(fileContent, &resume)
	if err != nil {
		writeJSONError(w, "Invalid file format", http.StatusBadRequest)
		return
	}

	err = ValidateResume(resume)
	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "ok"})
}

func writeJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}
