package single

import (
	"github.com/gin-gonic/gin"
	. "light-iot/base/domain"
	"light-iot/base/utils"
	. "light-iot/device/api"
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
