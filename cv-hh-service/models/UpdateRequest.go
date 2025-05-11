package models

type WorkExperience struct {
	Company     string `json:"company"`
	CompanyUrl  string `json:"company_url"`
	Position    string `json:"position"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Description string `json:"description"`
}

type ResumeUpdateRequest struct {
	SkillSet   []string         `json:"skill_set,omitempty"`
	Experience []WorkExperience `json:"experience,omitempty"`
}
