package main

import (
	"hui-iot/iot-worker/device_event"
	"hui-iot/iot-worker/iotevent"
	"log"
)

func main() {
	iotEvent := iotevent.IotEvent{}
	log.Printf("启动iot-worker服务成功！")
	iotEvent.AddOperation(device_event.ReadAttrEventOperation).
		Run("tcp://127.0.0.1:1883", []byte("iot-worker"))
}
