package event_decode

import (
	"encoding/json"
	"hui-iot/iot-driver/common"
	"log"
	"strings"
)

//温度过高事件
var (
	TempOverEventOperations = &common.Operation{Context: &common.Context{
		FromContext: common.FromContext{
			FromTopic: "/from/+/+/device/read_attr/up",
			FromQos:   0,
		},
		Ip:        "",
		UserName:  "",
		CilentID:  "",
		Payload:   nil,
		IsPublish: false,
	}, CodeHandler: decode}
)

func decode(cx *common.Context) *common.Context {
	var data map[string]interface{}
	json.Unmarshal(cx.Payload, &data)
	if temp := data["temp"].(float64); temp > 100 {
		cx.IsPublish = true
		data["temp"] = 1020
		str, _ := json.Marshal(&data)

		strs := strings.Split(cx.Topic, "/")
		toCx := common.ToContext{
			ToTopic:   "/sys/" + strs[2] + "/" + strs[3] + "/device/event/post",
			ToQos:     0,
			ToPayload: str,
		}
		cx.ToContexts[0] = toCx
	}
	log.Printf("paload:%s", cx.Payload)
	return cx
}
