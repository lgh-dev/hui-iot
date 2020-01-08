package api

import (
	"github.com/gin-gonic/gin"
	"light-iot/device/api/single"
	//batch "light-iot/device/api/batch"
)

func StartServer() {
	r := gin.Default()
	r.POST(single.PATH_DEVICE_ADD, single.DeviceAdd)
	r.POST(single.PATH_DEVICE_UPDATE, single.DeviceUpdate)
	r.POST(single.PATH_DEVICE_DELETE, single.DeviceDelete)
	r.POST(single.PATH_DEVICE_QUERY, single.DeviceQuery)
	r.Run(":9000")
}
