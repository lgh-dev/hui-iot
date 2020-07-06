package main

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc: Service of iot-iot-server's run file
 */
import (
	"github.com/prometheus/common/log"
	"hui-iot/iot-server/api"
	"hui-iot/iot-server/config"
)

func main() {
	log.Info("iot-server start!")
	port := config.Conf.GetString("server.port")
	api.GetServer().Run(":" + port)
}
