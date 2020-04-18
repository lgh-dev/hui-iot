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
	"hui-iot/iot-server/devicemodel/api"
)

func GetServer() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1/")

	v1.GET("devicemodels", api.FindAllDeviceModels) //find all

	v1.GET("device", nil)                          //find by ID
	v1.GET("devices", nil)                         //find by Page
	v1.POST("device", nil)                         //add
	v1.DELETE("device", nil)                       //delete
	v1.PUT("device", nil)                          //update
	v1.POST("device/:deviceId/command/:uKey", nil) //send cmd

	v1.GET("device/:deviceId/sensors", nil)  //find by date and deviceId and data number
	v1.GET("device/:deviceId/alarms", nil)   //find by date and deviceId and data number
	v1.GET("device/:deviceId/commands", nil) //find by date and deviceId and data number

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
