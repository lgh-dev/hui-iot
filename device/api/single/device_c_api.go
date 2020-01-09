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
	PATH_DEVICE_C        = "/device/c"
	PATH_DEVICE_C_QUERY  = utils.PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_C + utils.PATH_QUERY
	PATH_DEVICE_C_UPDATE = utils.PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_C + utils.PATH_UPDATE
)

func DeviceCUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}

func DeviceCQuery(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}
