package dependencies

type Project struct {
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	Version      string            `json:"version"`
	Homepage     string            `json:"homepage"`
	MainSrc      string            `json:"mainsrc"`
	Projects     []string          `json:"projects"`
	Scripts      map[string]string `json:"scripts,omitempty"`
	Dependencies map[string]string `json:"dependencies"`
	Locked       ProjectLock       `json:"-"`
}

func MakeNew() *Project {
	project := &Project{
		Dependencies: make(map[string]string),
		Projects:     []string{},
	}

	// project.Locked = LoadProjectLock(project)

	return project
}
