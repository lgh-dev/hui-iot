package api

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/gin-gonic/gin"
	. "hui-iot/iot-server/common"
	"net/http"
)

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/4/5 23:04
 * @Desc:
 */
func FindAllDeviceModels(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &DeviceModelMap)
}
