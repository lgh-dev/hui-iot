package single

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/gin-gonic/gin"
	. "light-iot/base/domain"
	"light-iot/base/utils"
	"net/http"
)

const (
	PATH_DEVICE_F        = "/device/f"
	PATH_DEVICE_F_UPDATE = utils.PATH_IOT_DEVICE_BATCH + PATH_DEVICE_F + utils.PATH_UPDATE
	PATH_DEVICE_F_EXEC   = utils.PATH_IOT_DEVICE_BATCH + PATH_DEVICE_F + utils.PATH_EXEC
)

func DeviceFuncExec(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}

func DeviceFuncUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}
