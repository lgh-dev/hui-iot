package single

import (
	"github.com/gin-gonic/gin"
	. "light-iot/base/domain"
	"light-iot/base/utils"
	. "light-iot/device/api"
	"net/http"
)

const (
	PATH_DEVICETYPE       = "/devicetype"
	PATH_DEVICETYPE_QUERY = utils.PATH_IOT_DEVICE_BATCH + PATH_DEVICETYPE + utils.PATH_ADD
)

func DeviceTypeQuery(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}
