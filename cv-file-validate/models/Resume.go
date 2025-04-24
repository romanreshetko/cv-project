package models

type Resume struct {
	Name               string      `yaml:"name"`
	Initials           string      `yaml:"initials"`
	Location           string      `yaml:"location"`
	LocationLink       string      `yaml:"locationLink"`
	AvatarUrl          string      `yaml:"avatarUrl"`
	PersonalWebsiteUrl string      `yaml:"personalWebsiteUrl,omitempty"`
	Summary            string      `yaml:"summary"`
	About              string      `yaml:"about"`
	Contact            Contact     `yaml:"contact"`
	Skills             []string    `yaml:"skills"`
	Work               []Work      `yaml:"work"`
	Education          []Education `yaml:"education,omitempty"`
	Projects           []Project   `yaml:"projects,omitempty"`
}

type Contact struct {
	Email  string   `yaml:"email"`
	Tel    string   `yaml:"tel"`
	Social []Social `yaml:"social,omitempty"`
}

type Social struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

type Education struct {
	School string `yaml:"school"`
	Degree string `yaml:"degree"`
	Start  string `yaml:"start"`
	End    string `yaml:"end"`
}

type Work struct {
	Company     string   `yaml:"company"`
	Link        string   `yaml:"link"`
	Badge       []string `yaml:"badge"`
	Title       string   `yaml:"title"`
	Start       string   `yaml:"start"`
	End         string   `yaml:"end"`
	Description string   `yaml:"description"`
}

type Project struct {
	Title       string   `yaml:"title"`
	Link        Link     `yaml:"link"`
	TechStack   []string `yaml:"techStack"`
	Description string   `yaml:"description"`
}

type Link struct {
	Label string `yaml:"label"`
	Href  string `yaml:"href"`
}
