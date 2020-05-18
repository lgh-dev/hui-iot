package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDeviceModels(t *testing.T) {
	ok := true
	if len(DeviceModelMap) == 0 {
		ok = false
	}
	for key := range DeviceModelMap {
		if len(DeviceModelMap[key]) == 0 {
			ok = false
		}
	}
	assert.True(t, ok)
}
