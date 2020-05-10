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
	FromTopic string
	Topic     string
	ToTopic   string
	FromQos   byte
	toQos     byte
	Ip        string
	UserName  string
	CilentID  string
	Payload   []byte
	IsPublish bool
}
type CodeHandler func(*Context) *Context

func (driver *Driver) AddOperation(operation *Operation) *Driver {
	driver.Operations = append(driver.Operations, operation)
	return driver
}

func (driver *Driver) Run(mqttServerAddress string, clientId []byte) {

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
				client.Publish(cx.ToTopic, cx.toQos, false, cx.Payload)
			}
			waitGroup.Add(1)
			go func(operation *Operation) {
				for {
					i++
					time.Sleep(time.Duration(3) * time.Second)
					token := client.Subscribe(operation.Context.FromTopic, operation.Context.FromQos, messageHandler)
					token.Wait()
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
