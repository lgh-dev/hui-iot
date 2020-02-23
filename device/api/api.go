package api

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"hui-iot/device/api/single"
)

func GetServer() *gin.Engine {
	r := gin.Default()
	//v1 := r.Group("/api/svc-iot-device/v1")
	//{
	//single interface
	//v1.GET("/devicetype/:id", single.UserGet)
	//v1.GET("/devicetype/:id", single.DeviceTypeQuery)
	//v1.GET("/devicetype/", single.DeviceTypeQuery)

	//v1.POST("/device/:id", single.GetDevice)
	//v1.POST("/device", single.AddDevice)
	//v1.PUT("/device", single.UpdateDevice)
	//v1.DELETE("/device", single.DeleteDevice)
	//
	//v1.GET("/device/:id/invariant-attr/:attrId", single.GetDeviceInvarintAttr)
	//v1.POST("/device/:id/invariant-attr/", single.UpdateDeviceInvarintAttr)
	//
	//v1.GET("/device/:id/read-attr/:attrId", single.GetDeviceReadAttr)
	//v1.POST("/device/:id/read-attr/", single.UpdateDeviceReadAttr)
	//
	//v1.GET("/device/:id/config-attr/:attrId", single.GetDeviceConfigAttr)
	//v1.POST("/device/:id/config-attr/", single.GetDeviceConfigAttr)
	//
	//v1.GET("/device/:id/func/:funcName", single.ExecDeviceFunc)
	//}
	r.GET("/devicetype/", single.UserGet)

	// batch interface
	//r.POST(PATH_DEVICE_BATCH_ADD, batch.DeviceAdd)
	//r.POST(PATH_DEVICE_BATCH_UPDATE, batch.DeviceUpdate)
	//r.POST(PATH_DEVICE_BATCH_DELETE, batch.DeviceDelete)
	//r.POST(PATH_DEVICE_BATCH_QUERY, batch.DeviceQuery)
	//
	//r.POST(PATH_DEVICE_BATCH_P_QUERY, batch.DevicePQuery)
	//r.POST(PATH_DEVICE_BATCH_P_UPDATE, batch.DevicePUpdate)
	//
	//r.POST(PATH_DEVICE_BATCH_R_QUERY, batch.DeviceRQuery)
	//r.POST(PATH_DEVICE_BATCH_R_UPDATE, batch.DeviceRUpdate)
	//
	//r.POST(PATH_DEVICE_BATCH_C_QUERY, batch.DeviceConfigQuery)
	//
	//r.POST(PATH_DEVICE_BATCH_F_UPDATE, batch.DeviceFUpdate)
	//r.POST(PATH_DEVICE_BATCH_F_EXEC, batch.DeviceFExec)
	//swagger2.0
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
