package iotevent

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"runtime"
	"sync"
	"time"
)

// 驱动
type IotEvent struct {
	Operations []*Operation
}

type Operation struct {
	Context      *Context
	EventHandler EventHandler
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
	//获取同步器
	waitGroup := sync.WaitGroup{}
	defer waitGroup.Done()
	//获取MQTT客戶端
	client := getClient(mqttServerAddress, clientId)

	//启动事件监听
	if iotEvent.Operations == nil || len(iotEvent.Operations) == 0 {
		log.Panic("不存在事件监听处理器！")
	}

	for _, operation := range iotEvent.Operations {
		//客户端连接判断
		if token := client.Connect(); token.WaitTimeout(time.Duration(60)*time.Second) && token.Wait() && token.Error() != nil {
			log.Printf("[Sub] mqtt connect error, taskId: %d, fail_nums: %d, error: %s \n", clientId, 1, token.Error())
			return
		}
		//获取消息处理器
		messageHandler := getMessageHandler(operation)
		waitGroup.Add(1)
		go func(operation *Operation) {
			for {
				time.Sleep(time.Duration(10) * time.Millisecond)
				token := client.Subscribe(operation.Context.FromTopic, operation.Context.FromQos, messageHandler)
				token.Wait()
			}
		}(operation)
		for {
			time.Sleep(time.Duration(1) * time.Second)
			log.Printf("当前协程数量：%d", runtime.NumGoroutine())
		}
		client.Disconnect(250)
		log.Println("[Sub] task is ok")
	}
}

func getMessageHandler(operation *Operation) mqtt.MessageHandler {
	var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		operation.Context.Topic = msg.Topic()
		operation.Context.Payload = msg.Payload()
		cx := operation.Context
		cx = operation.EventHandler(cx)
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
	clinetOptions.SetConnectTimeout(time.Duration(60) * time.Second)
	//创建客户端连接
	client := mqtt.NewClient(clinetOptions)
	return client
}
