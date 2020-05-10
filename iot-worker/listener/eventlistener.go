package listener

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	common2 "hui-iot/iot-server/common"
	"log"
	"strconv"
	//"database/sql"
	//_ "github.com/taosdata/TDengine/src/connector/go/src/taosSql"
)

// 接收事件
func ReceiveEvent(msg mqtt.Message) bool {
	log.Printf("msgId:" + strconv.Itoa(int(msg.MessageID())))
	return Handler(msg.Topic(), msg.Payload())
}

func Handler(topic string, payload []byte) bool {
	if topic == "" {
		logrus.Error("Receive event err of topic is nil")
		return false
	}
	//校验消息体
	if !gjson.ValidBytes(payload) {
		logrus.Errorf("Receive event's payload isn't Json : %s", string(payload))
		return false
	}
	//获取型号ID和设备ID。
	deviceModeID, deviceID := common2.GetDeviceModelIDAndDeviceIDForTopic(topic)
	fmt.Printf("deviceModelID:%s,deviceID:%s\n", deviceModeID, deviceID)

	//获取消息内容。
	log.Printf("read attr %s", payload)
	params := gjson.GetBytes(payload, "params")
	if params.Exists() {
		params.ForEach(func(key, value gjson.Result) bool {
			fmt.Printf("key:%s,value:%s", key, value.Str)
			return true
		})
	}
	return true
}

func converData() string {
	return ""
}

func saveData() {

}
