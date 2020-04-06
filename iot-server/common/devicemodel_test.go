package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDeviceModels(t *testing.T) {
	assert.True(t, len(DeviceModelMap) > 0)
}
