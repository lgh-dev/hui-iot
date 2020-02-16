package main

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc: Service of iot-device's run file
 */
import (
	"hui-iot/device/api"
	_ "hui-iot/device/docs"
)

// @title Go-site Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1
// @BasePath ""
func main() {
	r := api.GetServer()
	r.Run(":9000")

}
