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

// @title gin 框架
// @version 1.0
// @description 给予gin web框架搭建的业务骨架
// @termsofservice http://swagger.io/terms/
// @contact.name jinchunguang
// @contact.email jin-chunguang@foxmail.com
// @host localhost:10010"
func main() {
	r := api.GetServer()
	r.Run(":9000")

}
