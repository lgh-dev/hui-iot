package main

import (
	"hui-iot/iot-worker/config"
	"hui-iot/iot-worker/device_event"
	"hui-iot/iot-worker/iotevent"
	"log"
)

func main() {
	serviceAddr := config.Conf.GetString("mqtt.emqx.service_addr")
	clientId := config.Conf.GetString("mqtt.emqx.client_id")
	iotEvent := iotevent.IotEvent{}
	log.Printf("启动iot-worker服务成功，连接mqtt服务地址[%s]，客户端ID[%s]！\n", serviceAddr, clientId)
	iotEvent.AddOperation(device_event.ReadAttrEventOperation).
		Run(serviceAddr, []byte(clientId))
}
