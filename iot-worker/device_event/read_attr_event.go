package device_event

import (
	jsoniter "github.com/json-iterator/go"
	"hui-iot/iot-worker/iotevent"
	"hui-iot/iot-worker/listener"
	"hui-iot/iot-worker/utils"
	"log"
	"math"
	"time"
)

//请求Topic：/hiot/sys/{deviceModelID}/{deviceID}/dv/r/up
//响应Topic：/hiot/sys/{deviceModelID}/{deviceID}/dv/r/up_reponse
var ReadAttrEventOperation = &iotevent.Operation{Context: &iotevent.Context{
	FromContext: iotevent.FromContext{
		FromTopic: "/hiot/sys/+/+/dv/r/up",
		FromQos:   0,
	},
}, EventHandler: ReadAttrEventHandler, IsReceiveMessage: IsReceiveMessage}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var persistenceListener = listener.PersistenceListener{}

//判断是否接收该消息
func IsReceiveMessage(topic string, qos byte) bool {
	if topic[len(topic)-8:] == "/dv/r/up" {
		return true
	}
	return false
}

//topic:/hiot/sys/car/1111/dv/r/up
//payload:{"temp":34}
func ReadAttrEventHandler(c *iotevent.Context) *iotevent.Context {
	PrintMessage()
	deviceModel := utils.GetStrForTopic(c.Topic, 3)
	gatewayUID := utils.GetStrForTopic(c.Topic, 4)
	data := make(map[string]float64)
	err := json.Unmarshal(c.Payload, &data)
	if !utils.CheckErrorLn("解析json错误，只读数据格式：%s", err) {
		return c
	}
	//持久化事件
	eventDTO := listener.EventDTO{DeviceModel: &deviceModel, GatewayUID: &gatewayUID, EventType: listener.EventTypeRead, Event: &data}
	persistenceListener.DoEvent(&eventDTO)
	return c
}

var num = 0
var pow10 = int64(math.Pow10(9))
var startTime = time.Now().UnixNano()
var beforeTime = time.Now().UnixNano()

//处理方法
func PrintMessage() {
	num++
	if num%10000 == 0 {
		endTime := time.Now().UnixNano()
		var speed = 10000 * pow10 / (endTime - beforeTime)
		var avgSpeed = int64(num) * pow10 / (endTime - startTime)
		beforeTime = endTime
		log.Printf("累计消息数量：%d,消费速度：%d/s,平均速度:%d/s", num, speed, avgSpeed)
	}
}
