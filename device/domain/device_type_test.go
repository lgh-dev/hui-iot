package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDeviceTypes(t *testing.T) {
	GetDeviceTypes()
	assert.True(t, len(DeviceTypeMap) > 0, "Get DeviceTypes err")
}
