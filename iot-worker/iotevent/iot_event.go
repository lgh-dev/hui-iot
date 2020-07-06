package iotevent

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	libmqtt "github.com/goiiot/libmqtt"
	"hui-iot/iot-worker/utils"
	"log"
	"runtime"
	"time"
)

// 驱动
type IotEvent struct {
	Operations []*Operation
}

type Operation struct {
	Context          *Context
	EventHandler     EventHandler
	IsReceiveMessage func(topic string, qos byte) bool
}

//事件上下文
type Context struct {
	FromContext
	ToContext
	Topic    string
	Ip       string
	UserName string
	CilentID string
	Payload  []byte
}
type FromContext struct {
	FromTopic string
	FromQos   byte
}

type ToContext struct {
	ToTopic string
	toQos   byte
}
type EventHandler func(*Context) *Context

func (iotEvent *IotEvent) AddOperation(operation *Operation) *IotEvent {
	iotEvent.Operations = append(iotEvent.Operations, operation)
	return iotEvent
}

func (iotEvent *IotEvent) Run(mqttServerAddress string, clientId []byte) {
	//获取MQTT客戶端
	client := getClient("tcp://"+mqttServerAddress, clientId)

	//启动事件监听
	if iotEvent.Operations == nil || len(iotEvent.Operations) == 0 {
		log.Panic("不存在事件监听处理器！")
	}

	for _, operation := range iotEvent.Operations {
		//客户端连接判断
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			log.Printf("[Sub] mqtt connect error, taskId: %d, fail_nums: %d, error: %s \n", clientId, 1, token.Error())
			return
		}
		//获取消息处理器
		messageHandler := getMessageHandler(operation)
		Subscribe(client, messageHandler, clientId, operation)
	}
	for {
		time.Sleep(time.Duration(1) * time.Second)
		log.Printf("当前协程数量：%d", runtime.NumGoroutine())
	}
	client.Disconnect(250)
	log.Println("[Sub] task is ok")
}

func Subscribe(client mqtt.Client, messageHandler mqtt.MessageHandler, clientId []byte, operation *Operation) {
	func(operation *Operation) {
		token := client.Subscribe(operation.Context.FromTopic, operation.Context.FromQos, messageHandler)
		if token.Wait(); token.Error() != nil {
			log.Panicf("[Sub] mqtt connect error, taskId: %d, fail_nums: %d, error: %s \n", clientId, 1, token.Error())
		}
		log.Print("test subscribe")
	}(operation)
}

func getMessageHandler(operation *Operation) mqtt.MessageHandler {
	var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		operation.Context.Topic = msg.Topic()
		operation.Context.Payload = msg.Payload()
		cx := operation.Context
		operation.EventHandler(cx)
	}
	return messageHandler
}

//获取mqtt客户端
func getClient(mqttServerAddress string, clientId []byte) mqtt.Client {
	//设置连接参数
	clinetOptions := mqtt.NewClientOptions().AddBroker(mqttServerAddress)
	//设置客户端ID
	clinetOptions.SetClientID(string(clientId))
	//设置连接超时
	clinetOptions.SetConnectTimeout(time.Duration(360) * time.Second)
	clinetOptions.SetAutoReconnect(true)
	//创建客户端连接
	client := mqtt.NewClient(clinetOptions)
	return client
}

func (iotEvent *IotEvent) Run2(mqttServerAddress string, clientId []byte) {
	//获取MQTT客戶端
	client := getClient2(string(clientId))

	//启动事件监听
	if iotEvent.Operations == nil || len(iotEvent.Operations) == 0 {
		log.Panic("不存在事件监听处理器！")
	}

	topics := make([]*libmqtt.Topic, len(iotEvent.Operations))
	log.Println("Subscribe Topics:")
	for i, operation := range iotEvent.Operations {
		topics[i] = &libmqtt.Topic{Name: operation.Context.FromTopic, Qos: operation.Context.FromQos}
		log.Println(*topics[i])
		client.HandleTopic(".*", func(client libmqtt.Client, topic string, qos libmqtt.QosLevel, msg []byte) {
			operation.Context.Topic = topic
			operation.Context.Payload = msg
			cx := operation.Context
			operation.EventHandler(cx)
		})
	}
	client.Subscribe(topics...)

	err := client.ConnectServer(mqttServerAddress)
	utils.CheckErrorLn("connect  MQTT server err %s,server addr [%s]\n", err, mqttServerAddress)
	client.Wait()
}

//获取mqtt客户端
func getClient2(clientID string) libmqtt.Client {
	client, err := libmqtt.NewClient(
		libmqtt.WithVersion(libmqtt.V311, true),
		// enable keepalive (10s interval) with 20% tolerance
		libmqtt.WithKeepalive(10, 1.2),
		// enable auto reconnect and set backoff strategy
		libmqtt.WithAutoReconnect(true),
		libmqtt.WithBackoffStrategy(time.Second, 500*time.Second, 1.2),
		// use RegexRouter for topic routing if not specified
		// will use TextRouter, which will match full text
		libmqtt.WithRouter(libmqtt.NewRegexRouter()),
		libmqtt.WithClientID(clientID),
		libmqtt.WithUnsubHandleFunc(func(client libmqtt.Client, topics []string, err error) {
			log.Println("取消了订阅!")
		}),
	)

	if err != nil {
		// handle client creation error
		panic("create mqtt client failed")
	}
	return client
}
