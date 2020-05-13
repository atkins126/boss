package machineinfo

import (
	"crypto/md5"
	"encoding/hex"
	"io"

	"github.com/denisbrodbeck/machineid"
	log "github.com/sirupsen/logrus"
)

const defaultID = "459fd21a-31cf-57b9-b6a7-1625ccad3d76"

var (
	MachineID     = machineID()
	Md5MachineID  = md5MachineID()
	MachineIDByte = []byte(MachineID[:16])
)

func machineID() string {
	id, err := machineid.ID()
	if err != nil {
		log.Error("Error on get machine ID, assuming default.")
		id = defaultID
	}
	return id
}

func md5MachineID() string {
	hasher := md5.New()
	if _, err := io.WriteString(hasher, machineID()); err != nil {
		log.Warn("Failed on  write machine id to hash")
	}
	return hex.EncodeToString(hasher.Sum(nil))
}
