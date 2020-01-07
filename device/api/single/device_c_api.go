package single

import (
	"github.com/gin-gonic/gin"
	. "light-iot/base/domain"
	"light-iot/base/utils"
	"net/http"
)

const (
	PATH_DEVICE_C       = "/device/c"
	PATH_DEVICE_C_QUERY = utils.PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_F + utils.PATH_QUERY
)

func DeviceConfigUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}

func DeviceConfigQuery(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}
