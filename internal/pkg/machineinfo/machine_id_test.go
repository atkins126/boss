package machineinfo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMachineID(t *testing.T) {
	assert.Equal(t, machineID(), MachineID)
	assert.Equal(t, md5MachineID(), Md5MachineID)
}
