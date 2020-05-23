package common

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"sync"
	"time"
)

// 驱动
type Driver struct {
	Operations []*Operation
}

type Operation struct {
	Context     *Context
	CodeHandler CodeHandler
}

//事件上下文
type Context struct {
	FromContext
	ToContexts []ToContext
	Topic      string
	Ip         string
	UserName   string
	CilentID   string
	Payload    []byte
	IsPublish  bool
}
type FromContext struct {
	FromTopic string
	FromQos   byte
}

type ToContext struct {
	ToTopic   string
	ToQos     byte
	ToPayload []byte
}

type CodeHandler func(*Context) *Context

func (driver *Driver) AddOperation(operation *Operation) *Driver {
	driver.Operations = append(driver.Operations, operation)
	return driver
}

func (driver *Driver) Run(mqttServerAddress string, clientId []byte) {

	waitGroup := sync.WaitGroup{}
	defer waitGroup.Done()
	client := GetMQTTClient(mqttServerAddress, clientId)

	//启动事件监听
	if driver.Operations != nil {
		for _, operation := range driver.Operations {
			//客户端连接判断
			if token := client.Connect(); token.WaitTimeout(time.Duration(60)*time.Second) && token.Wait() && token.Error() != nil {
				log.Printf("[Sub] mqtt connect error, taskId: %d, fail_nums: %d, error: %s \n", clientId, 1, token.Error())
				return
			}
			i := 0
			//define a function for the default message handler
			var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
				operation.Context.Topic = msg.Topic()
				operation.Context.Payload = msg.Payload()
				cx := operation.Context
				cx = operation.CodeHandler(cx)
				if cx.ToContexts != nil {
					for _, toCx := range cx.ToContexts {
						client.Publish(toCx.ToTopic, toCx.ToQos, false, toCx.ToPayload)
					}
				}
			}
			waitGroup.Add(1)
			go func(operation *Operation) {
				token := client.Subscribe(operation.Context.FromTopic, operation.Context.FromQos, messageHandler)
				token.Wait()
				if token.Error() != nil {
					log.Panicf("[Sub] mqtt connect error, taskId: %d, fail_nums: %d, error: %s \n", clientId, 1, token.Error())
				}
			}(operation)
			for {
				i++
				time.Sleep(time.Duration(3) * time.Second)
			}
			client.Disconnect(250)
			log.Println("[Sub] task is ok")

		}
	}
}

func GetMQTTClient(mqttServerAddress string, clientId []byte) mqtt.Client {
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
