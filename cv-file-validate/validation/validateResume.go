package validation

import (
	. "cv-file-validate/models"
	"errors"
	"fmt"
)

func ValidateResume(resume Resume) error {
	var validationErrors []string

	if resume.Name == "" {
		validationErrors = append(validationErrors, "name")
	}
	if resume.Initials == "" {
		validationErrors = append(validationErrors, "initials")
	}
	if resume.Location == "" {
		validationErrors = append(validationErrors, "location")
	}
	if resume.LocationLink == "" {
		validationErrors = append(validationErrors, "location_link")
	}
	if resume.AvatarUrl == "" {
		validationErrors = append(validationErrors, "avatar_url")
	}
	if resume.Summary == "" {
		validationErrors = append(validationErrors, "summary")
	}
	if resume.About == "" {
		validationErrors = append(validationErrors, "about")
	}
	if resume.Contact.Email == "" {
		validationErrors = append(validationErrors, "contact.email")
	}
	if resume.Contact.Tel == "" {
		validationErrors = append(validationErrors, "contact.tel")
	}
	for i, edu := range resume.Education {
		if edu.School == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("education[%d].school", i))
		}
		if edu.Degree == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("education[%d].degree", i))
		}
		if edu.Start == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("education[%d].start", i))
		}
		if edu.End == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("education[%d].end", i))
		}
	}
	if len(resume.Work) == 0 {
		validationErrors = append(validationErrors, "work")
	}
	for i, job := range resume.Work {
		if job.Company == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("work[%d].company", i))
		}
		if job.Title == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("work[%d].title", i))
		}
		if job.Start == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("work[%d].start", i))
		}
		if job.Description == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("work[%d].description", i))
		}
	}
	if len(resume.Skills) == 0 {
		validationErrors = append(validationErrors, "skills")
	}
	for i, project := range resume.Projects {
		if project.Title == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("projects[%d].title", i))
		}
		if len(project.TechStack) == 0 {
			validationErrors = append(validationErrors, fmt.Sprintf("projects[%d].techStack", i))
		}
		if project.Description == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("projects[%d].description", i))
		}
	}

	if len(validationErrors) > 0 {
		return errors.New("validation errors: missing required fields\n" + fmt.Sprintf("%v", validationErrors))
	}

	return nil
}

func IsValidYAML(filename string) bool {
	return len(filename) > 4 && (filename[len(filename)-5:] == ".yaml" || filename[len(filename)-4:] == ".yml")
}
