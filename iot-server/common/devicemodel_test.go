package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDeviceModels(t *testing.T) {
	// config file path
	const configPath = "../"
	isOK, deviceModelMap := GetDeviceModels(configPath)
	assert.True(t, isOK, "Get DeviceModels err")
	assert.True(t, len(deviceModelMap) > 0)
}
