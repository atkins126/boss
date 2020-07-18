package bosspackage

type BossPackage struct {
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	Version      string            `json:"version"`
	Homepage     string            `json:"homepage"`
	MainSrc      string            `json:"mainsrc"`
	Projects     []string          `json:"projects"`
	Scripts      map[string]string `json:"scripts,omitempty"`
	Dependencies map[string]string `json:"dependencies"`
	Locked       *ProjectLock      `json:"-"`
}

func MakeNew() *BossPackage {
	project := BossPackage{
		Dependencies: make(map[string]string),
		Projects:     []string{},
		Locked:       &ProjectLock{},
	}

	return &project
}

func LoadOrNew(path string) *BossPackage {
	project := MakeNew()

	project.Locked = LoadProjectLock(project)

	return project
}
