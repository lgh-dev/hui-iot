package api

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/gin-gonic/gin"
	batch "light-iot/device/api/batch"
	"light-iot/device/api/single"
)

func StartServer() {
	r := gin.Default()
	//single interface
	r.POST(single.PATH_DEVICE_ADD, single.DeviceAdd)
	r.POST(single.PATH_DEVICE_UPDATE, single.DeviceUpdate)
	r.POST(single.PATH_DEVICE_DELETE, single.DeviceDelete)
	r.POST(single.PATH_DEVICE_QUERY, single.DeviceQuery)

	r.POST(single.PATH_DEVICE_P_QUERY, single.DevicePQuery)
	r.POST(single.PATH_DEVICE_P_UPDATE, single.DevicePUpdate)

	r.POST(single.PATH_DEVICE_R_QUERY, single.DeviceRQuery)
	r.POST(single.PATH_DEVICE_R_UPDATE, single.DeviceRUpdate)

	r.POST(single.PATH_DEVICE_C_QUERY, single.DeviceCQuery)
	r.POST(single.PATH_DEVICE_C_UPDATE, single.DeviceCUpdate)

	r.POST(single.PATH_DEVICE_F_UPDATE, single.DeviceFUpdate)
	r.POST(single.PATH_DEVICE_F_EXEC, single.DeviceFExec)

	// batch interface
	r.POST(batch.PATH_DEVICE_ADD, batch.DeviceAdd)
	r.POST(batch.PATH_DEVICE_UPDATE, batch.DeviceUpdate)
	r.POST(batch.PATH_DEVICE_DELETE, batch.DeviceDelete)
	r.POST(batch.PATH_DEVICE_QUERY, batch.DeviceQuery)

	r.POST(batch.PATH_DEVICE_P_QUERY, batch.DevicePQuery)
	r.POST(batch.PATH_DEVICE_P_UPDATE, batch.DevicePUpdate)

	r.POST(batch.PATH_DEVICE_R_QUERY, batch.DeviceRQuery)
	r.POST(batch.PATH_DEVICE_R_UPDATE, batch.DeviceRUpdate)

	r.POST(batch.PATH_DEVICE_C_QUERY, batch.DeviceConfigQuery)

	r.POST(batch.PATH_DEVICE_F_UPDATE, batch.DeviceFUpdate)
	r.POST(batch.PATH_DEVICE_F_EXEC, batch.DeviceFExec)

	r.Run(":9000")
}
