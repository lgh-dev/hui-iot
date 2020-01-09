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
	PATH_DEVICE_R        = "/device/r"
	PATH_DEVICE_R_UPDATE = utils.PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_R + utils.PATH_UPDATE
	PATH_DEVICE_R_QUERY  = utils.PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_R + utils.PATH_QUERY
)

func DeviceRUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}

func DeviceRQuery(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}
