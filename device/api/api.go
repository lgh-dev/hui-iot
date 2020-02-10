package api

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/gin-gonic/gin"
	batch "hui-iot/device/api/batch"
	single "hui-iot/device/api/single"
)

func StartServer() *gin.Engine {
	r := gin.Default()
	//single interface
	r.POST(PATH_DEVICETYPE_SINGLE_QUERY, single.DeviceTypeQuery)

	r.POST(PATH_DEVICE_SINGLE_ADD, single.DeviceAdd)
	r.POST(PATH_DEVICE_SINGLE_UPDATE, single.DeviceUpdate)
	r.POST(PATH_DEVICE_SINGLE_DELETE, single.DeviceDelete)
	r.POST(PATH_DEVICE_SINGLE_QUERY, single.DeviceQuery)

	r.POST(PATH_DEVICE_SINGLE_P_QUERY, single.DevicePQuery)
	r.POST(PATH_DEVICE_SINGLE_P_UPDATE, single.DevicePUpdate)

	r.POST(PATH_DEVICE_SINGLE_R_QUERY, single.DeviceRQuery)
	r.POST(PATH_DEVICE_SINGLE_R_UPDATE, single.DeviceRUpdate)

	r.POST(PATH_DEVICE_SINGLE_C_QUERY, single.DeviceCQuery)
	r.POST(PATH_DEVICE_SINGLE_C_UPDATE, single.DeviceCUpdate)

	r.POST(PATH_DEVICE_SINGLE_F_UPDATE, single.DeviceFUpdate)
	r.POST(PATH_DEVICE_SINGLE_F_EXEC, single.DeviceFExec)

	// batch interface
	r.POST(PATH_DEVICE_BATCH_ADD, batch.DeviceAdd)
	r.POST(PATH_DEVICE_BATCH_UPDATE, batch.DeviceUpdate)
	r.POST(PATH_DEVICE_BATCH_DELETE, batch.DeviceDelete)
	r.POST(PATH_DEVICE_BATCH_QUERY, batch.DeviceQuery)

	r.POST(PATH_DEVICE_BATCH_P_QUERY, batch.DevicePQuery)
	r.POST(PATH_DEVICE_BATCH_P_UPDATE, batch.DevicePUpdate)

	r.POST(PATH_DEVICE_BATCH_R_QUERY, batch.DeviceRQuery)
	r.POST(PATH_DEVICE_BATCH_R_UPDATE, batch.DeviceRUpdate)

	r.POST(PATH_DEVICE_BATCH_C_QUERY, batch.DeviceConfigQuery)

	r.POST(PATH_DEVICE_BATCH_F_UPDATE, batch.DeviceFUpdate)
	r.POST(PATH_DEVICE_BATCH_F_EXEC, batch.DeviceFExec)
	return r
}
