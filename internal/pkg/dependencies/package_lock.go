package dependencies

import "time"

type DependencyArtifacts struct {
	Bin []string `json:"bin,omitempty"`
	Dcp []string `json:"dcp,omitempty"`
	Dcu []string `json:"dcu,omitempty"`
	Bpl []string `json:"bpl,omitempty"`
}

type LockedDependency struct {
	Name      string              `json:"name"`
	Version   string              `json:"version"`
	Hash      string              `json:"hash"`
	Artifacts DependencyArtifacts `json:"artifacts"`
	Failed    bool                `json:"failed"`
	Changed   bool                `json:"changed"`
}

type ProjectLock struct {
	Hash      string                      `json:"hash"`
	Created   time.Time                   `json:"createdAt"`
	Installed map[string]LockedDependency `json:"installedModules"`
}

func LoadProjectLock(project *ProjectLock) *ProjectLock {
	return &ProjectLock{}
}
