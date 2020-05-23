package main

import (
	"flag"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/thinkeridea/go-extend/exstrings"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	clientNumber = uint64(10000) //客户端数量
	interval     = 1             //秒
)

func main() {
	//生成连接的客户端数
	c := flag.Uint64("n", clientNumber, "client nums")
	flag.Parse()
	nums := int(*c)
	wg := sync.WaitGroup{}
	for i := 0; i < nums; i++ {
		wg.Add(1)
		time.Sleep(5 * time.Millisecond)
		go createTask(i, &wg)
	}
	wg.Wait()
}

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
var fail_nums int = 0

func createTask(taskId int, wg *sync.WaitGroup) {
	defer wg.Done()
	//opts := MQTT.NewClientOptions().AddBroker("tcp://192.168.20.101:1883").SetUsername("test").SetPassword("test")
	opts := MQTT.NewClientOptions().AddBroker("tcp://127.0.0.1:1883").SetUsername("test").SetPassword("test")
	opts.SetClientID(fmt.Sprintf("go-simple-client:%d-%d", taskId, time.Now().Unix()))
	opts.SetDefaultPublishHandler(f)
	opts.SetConnectTimeout(time.Duration(60) * time.Second)

	//创建连接
	c := MQTT.NewClient(opts)

	if token := c.Connect(); token.WaitTimeout(time.Duration(60)*time.Second) && token.Wait() && token.Error() != nil {
		fail_nums++
		log.Printf("taskId:%d,fail_nums:%d,error:%s \n", taskId, fail_nums, token.Error())
		return
	}

	//每隔5秒向topic发送一条消息
	i := 0
	gatewayUID := strconv.Itoa(taskId)
	topic := "/hiot/sys/smart_car_camera/?/dv/r/up"
	topic = exstrings.Replace(topic, "?", gatewayUID, 1)
	payload := "{\"temp\":?,\"msgId\":?}"
	payload = exstrings.Replace(payload, "?", strconv.Itoa(rand.Intn(100)), 1)
	payload = exstrings.Replace(payload, "?", strconv.Itoa(int(rand.Int31())), 1)
	for {
		i++
		time.Sleep(time.Duration(interval) * time.Second)
		token := c.Publish(topic, 0, false, payload)
		token.Wait()
	}

	c.Disconnect(250)
	fmt.Println("task ok!!")
}
