package main

import (
	"hui-iot/iot-driver/common"
	event_decode "hui-iot/iot-driver/temperature-sensor/event-decode"
)

func main() {
	dr := common.Driver{}
	dr.AddOperation(event_decode.TempOverEventOperations)
	dr.Run("tcp://127.0.0.1:1883", []byte("driver-temperature-sensor"))
}
