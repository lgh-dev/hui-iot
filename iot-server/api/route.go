package api

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/go-playground/validator.v9"
	. "hui-iot/iot-server/controller"
	"hui-iot/iot-server/middleware"
	"hui-iot/iot-server/valid"
	"net/http"
)

var applicationApi ApplicationApi

var deviceApi DeviceApi

var deviceModelApi DeviceModelApi

func GetServer() *gin.Engine {
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//gin.DefaultErrorWriter
	//= io.MultiWriter(f)
	//设置生产模式，性能提升一倍。
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	//r.Use(gin.Logger(), gin.Recovery(), middleware.AppAuth())
	//不打日志性能差6~8倍。
	r.Use(gin.Recovery(), middleware.AppAuth())
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("checkDeviceNameOnly", valid.CheckDeviceNameOnly, false)
	}
	//前端资源
	r.StaticFS("iot-frontend", http.Dir("../iot-frontend"))

	v1 := r.Group("/api/v1/")

	// device model api
	v1.GET("devicemodels", deviceModelApi.FindAll)     //find all
	v1.GET("devicemodel/:id", deviceModelApi.FindById) //find by id
	// codec api
	v1.GET("codecs", nil)    //find all
	v1.GET("codec/:id", nil) //find by ID
	v1.GET("codec", nil)     //find by deviceModelId
	//app api
	v1.GET("app/jwt", applicationApi.GetToken) // get app token
	v1.GET("apps", nil)                        //find all

	// device api
	v1.POST("device", deviceApi.Add)                           //add [base attr & config attr & business attr]
	v1.GET("device/:id", deviceApi.FindById)                   //find by ID
	v1.GET("devices", deviceApi.FindByPage)                    //find [base attr|config attr|read attr|business attr] by Page
	v1.DELETE("device", deviceApi.Delete)                      //delete
	v1.PUT("device", deviceApi.Update)                         //update [base attr & config attr & business attr]
	v1.POST("device/:id/command/:uKey", deviceApi.SendCommand) //send cmd by ID、uKey and params use JSON to body
	// device data api
	v1.GET("device/:id/sensors", nil)  //find by date and deviceId and data number
	v1.GET("device/:id/alarms", nil)   //find by date and deviceId and data number
	v1.GET("device/:id/commands", nil) //find by date and deviceId and data number

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
