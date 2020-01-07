package single

import (
	"github.com/gin-gonic/gin"
	. "light-iot/base/domain"
	"light-iot/base/utils"
	. "light-iot/device/api"
	"net/http"
)

const (
	PATH_DEVICE_R        = "/device/p"
	PATH_DEVICE_R_UPDATE = utils.PATH_IOT_DEVICE_BATCH + PATH_DEVICE_R + utils.PATH_UPDATE
	PATH_DEVICE_R_QUERY  = utils.PATH_IOT_DEVICE_BATCH + PATH_DEVICE_R + utils.PATH_QUERY
)

func DeviceRUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}

func DeviceRQuery(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}
