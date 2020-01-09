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
	PATH_DEVICE_C       = "/device/c"
	PATH_DEVICE_C_QUERY = utils.PATH_IOT_DEVICE_BATCH + PATH_DEVICE_C + utils.PATH_QUERY
)

func DeviceConfigUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}

func DeviceConfigQuery(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}
