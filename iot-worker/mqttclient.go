package iot_worker

import (
	"fmt"
	"hui-iot/base/common"
	"sync"
	"time"

	//import the Paho Go MQTT library
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	mqttServerAddress = "tcp://192.168.20.101:1883"
	topic             = common.MatchingAll(common.TOPIC_READ_ATTR_POST)
	clientId          = []byte("svc-hui-iot-iot-worker")
)

//define a function for the default message handler
var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	ReceiveEvent(msg)
}

/***
*
* 连接任务和消息订阅方法
 */
func SubMsg() {
	waitGroup := sync.WaitGroup{}

	defer waitGroup.Done()
	//设置连接参数
	clinetOptions := mqtt.NewClientOptions().AddBroker(mqttServerAddress)
	//设置客户端ID
	clinetOptions.SetClientID(string(clientId))
	//设置连接超时
	clinetOptions.SetConnectTimeout(time.Duration(60) * time.Second)
	//创建客户端连接
	client := mqtt.NewClient(clinetOptions)

	//客户端连接判断
	if token := client.Connect(); token.WaitTimeout(time.Duration(60)*time.Second) && token.Wait() && token.Error() != nil {
		fmt.Printf("[Sub] mqtt connect error, taskId: %d, fail_nums: %d, error: %s \n", clientId, 1, token.Error())
		return
	}

	i := 0

	for {
		i++
		time.Sleep(time.Duration(3) * time.Second)
		//fmt.Printf("start publish msg to mqtt broker, taskId: %d, count: %d \n", taskId, i)
		//发布消息
		token := client.Subscribe(topic, 0, messageHandler)
		token.Wait()
	}

	client.Disconnect(250)
	fmt.Println("[Sub] task is ok")

}
