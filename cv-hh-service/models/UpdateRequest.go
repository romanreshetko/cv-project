package models

type WorkExperience struct {
	Company     string `yaml:"company"`
	CompanyUrl  string `yaml:"company_url"`
	Position    string `yaml:"position"`
	Start       string `yaml:"start"`
	End         string `yaml:"end"`
	Description string `yaml:"description"`
}

type ResumeUpdateRequest struct {
	SkillSet   []string         `json:"skill_set,omitempty"`
	Experience []WorkExperience `json:"experience,omitempty"`
}
