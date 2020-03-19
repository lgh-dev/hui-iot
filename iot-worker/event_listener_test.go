package iot_worker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	topic := "/sys/111/1111/thing/event/property/post"
	json := `
{
  "id": "123",
  "version": "1.0",
  "params": {
    "Power": {
      "value": "on",
      "time": 1524448722000
    },
    "WF": {
      "value": 23.6,
      "time": 1524448722000
    }
  },
  "method": "thing.event.property.post"
}
`
	assert.True(t, Handler(topic, []byte(json)), "Receive event err")
}
