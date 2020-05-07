package controller

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/gin-gonic/gin"
	"hui-iot/iot-server/config"
	"net/http"
)

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/4/5 23:04
 * @Desc:
 */
type DeviceModelApi struct {
}

func (deviceModelApi DeviceModelApi) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &config.DeviceModelMap)
}

func (deviceModelApi DeviceModelApi) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, config.DeviceModelMap[id])
}
