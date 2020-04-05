package api

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/gin-gonic/gin"
	"hui-iot/iot-server/common"
	"net/http"
	"sync"
)

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/4/5 23:04
 * @Desc:
 */
func FindAllDeviceModels(ctx *gin.Context) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	const configPath = "../"
	_, deviceModelMap := common.GetDeviceModels(configPath)
	ctx.JSON(http.StatusOK, common.BuildSucc(&common.ResultDTO{Data: deviceModelMap}))
}
