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
	. "hui-iot/iot-server/controller"
)

func GetServer() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1/")
	// device model api
	v1.GET("devicemodels", FindAllDeviceModels)    //find all
	v1.GET("devicemodel/:id", FindDeviceModelById) //find by id
	// codec api
	v1.GET("codecs", nil)    //find all
	v1.GET("codec/:id", nil) //find by ID
	v1.GET("codec", nil)     //find by deviceModelId
	//app api
	v1.GET("app/auth", nil) // get app token
	v1.GET("apps", nil)     //find all

	// device api
	v1.GET("device", nil)                    //find by ID
	v1.GET("devices", nil)                   //find [base attr|config attr|read attr|business attr] by Page
	v1.POST("device", nil)                   //add [base attr & config attr & business attr]
	v1.DELETE("device", nil)                 //delete
	v1.PUT("device", nil)                    //update [base attr & config attr & business attr]
	v1.POST("device/:id/command/:uKey", nil) //send cmd by ID、uKey and params use JSON to body
	// device data api
	v1.GET("device/:deviceId/sensors", nil)  //find by date and deviceId and data number
	v1.GET("device/:deviceId/alarms", nil)   //find by date and deviceId and data number
	v1.GET("device/:deviceId/commands", nil) //find by date and deviceId and data number

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
