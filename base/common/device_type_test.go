package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDeviceTypes(t *testing.T) {
	// config file path
	const configPath = "../config/devicetype/"
	assert.True(t, GetDeviceTypes(configPath), "Get DeviceTypes err")
}
