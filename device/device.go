package main

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc: Service of iot-device's run file
 */
import (
	"hui-iot/device/api"
)

func main() {
	r := api.StartServer()
	r.Run(":9000")

}
