package main

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc: Service of iot-iot-server's run file
 */
import (
	"hui-iot/iot-server/api"
	_ "hui-iot/iot-server/docs"
)

func main() {
	api.GetServer().Run(":9000")
}
