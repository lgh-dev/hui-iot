package utils

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetDeviceModelIDAndDeviceIDForTopic(t *testing.T) {
	topic := "/sys/111/1111/thing/event/property/post"
	deviceModeID := GetStrForTopic(topic, 2)
	deviceID := GetStrForTopic(topic, 3)
	assert.Equal(t, deviceModeID, "111", "Check deviceModeID is false")
	assert.Equal(t, deviceID, "1111", "Check deviceModeID is false")
}

//压测
func BenchmarkGetDeviceModelIDAndDeviceIDForTopic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		topic := "/sys/111/1111111/thing/event/property/post"
		GetStrForTopic(topic, 2)
		GetStrForTopic(topic, 3)
	}
}
