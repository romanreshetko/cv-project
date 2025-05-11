package handlers

import (
	"bytes"
	. "cv-hh-service/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gopkg.in/yaml.v3"
)

func UpdateResumeHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		writeJSONError(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("file")
	token := r.FormValue("token")
	resumeId := r.FormValue("resume_id")
	if err != nil {
		writeJSONError(w, "Failed to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

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

	updateResume := mapResumeToUpdateResume(resume)

	url := fmt.Sprintf("https://api.hh.ru/resumes/%s", resumeId)
	jsonBody, err := json.Marshal(updateResume)
	if err != nil {
		writeJSONError(w, "Failed to serialize resume", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		writeJSONError(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("HH-User-Agent", "Resume-As-Code")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		writeJSONError(w, "Failed to send request", http.StatusInternalServerError)
		return
	}
	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		writeJSONError(w, string(body), resp.StatusCode)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "ok"})
}

func mapResumeToUpdateResume(resume Resume) ResumeUpdateRequest {
	experience := make([]WorkExperience, 0, len(resume.Work))
	for _, exp := range resume.Work {
		work := WorkExperience{
			Company:     exp.Company,
			CompanyUrl:  exp.Link,
			Position:    exp.Title,
			Start:       exp.Start,
			End:         exp.End,
			Description: exp.Description,
		}
		experience = append(experience, work)
	}

	return ResumeUpdateRequest{
		SkillSet:   resume.Skills,
		Experience: experience,
	}
}
