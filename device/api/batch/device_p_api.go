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
	PATH_DEVICE_P        = "/device/p"
	PATH_DEVICE_P_UPDATE = utils.PATH_IOT_DEVICE_BATCH + PATH_DEVICE_P + utils.PATH_UPDATE
	PATH_DEVICE_P_QUERY  = utils.PATH_IOT_DEVICE_BATCH + PATH_DEVICE_P + utils.PATH_QUERY
)

func DevicePUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}

func DevicePQuery(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}
