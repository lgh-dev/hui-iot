package common

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetDeviceTypeIDAndDeviceIDForTopic(t *testing.T) {
	topic := "/sys/111/1111/thing/event/property/post"
	deviceModeID, deviceID := GetDeviceTypeIDAndDeviceIDForTopic(topic)
	assert.Equal(t, deviceModeID, "111", "Check deviceModeID is false")
	assert.Equal(t, deviceID, "1111", "Check deviceModeID is false")
}

//压测
func BenchmarkGetDeviceTypeIDAndDeviceIDForTopic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		topic := "/sys/111/1111111/thing/event/property/post"
		GetDeviceTypeIDAndDeviceIDForTopic(topic)
	}
}
