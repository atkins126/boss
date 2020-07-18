package cache

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/hashload/boss/internal/pkg/utils"
)

// Root
//   +-> Dependency 1
//   |     |-> sub dependency 2
//   |     |     +-> sub sub dependency 3
//   |     |           +-> None
//   |     |
//   |     +-> sub dependency 1
//   |           +-> None
//   |
//   +-> Dependency 2
//   |     +-> None
//    None

type StorageMetadata struct {
	path         string
	LastPurge    time.Time              `json:"last_purge"`
	Repositories map[string]*Repository `json:"repositories"`
}

type VersionInfo struct {
	Dependencies map[string]string `json:"dependencies"`
}

type Repository struct {
	Key             string                  `json:"key"`
	LastUpdate      time.Time               `json:"last_update"`
	LastUtilization time.Time               `json:"last_utilization"`
	Versions        map[string]*VersionInfo `json:"versions"`
}

func MakeStorage() *StorageMetadata {
	storage := &StorageMetadata{
		Repositories: make(map[string]*Repository),
	}
	return storage
}

func ReadMetadata(path string) *StorageMetadata {
	storage := MakeStorage()
	storage.path = path
	rawBytes, err := ioutil.ReadFile(path)
	if os.IsNotExist(err) {
		return storage
	}
	utils.CheckError(err)

	err = json.Unmarshal(rawBytes, storage)
	utils.CheckError(err)

	return storage
}

func (s *StorageMetadata) EnsureCache() {
	//TODO: save to file

}
