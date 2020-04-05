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
	//{
	//single interface
	v1.GET("devicemodels", api.FindAllDeviceModels)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
