package main

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc: Service of iot-iot-server's run file
 */
import (
	"github.com/sirupsen/logrus"
	"hui-iot/iot-server/api"
)

func main() {
	logrus.Info("iot-server started!")
	api.GetServer().Run(":9000")
}
