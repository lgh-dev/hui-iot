package single

import (
	"github.com/gin-gonic/gin"
	. "light-iot/base/domain"
	"light-iot/base/utils"
	"net/http"
)

const (
	PATH_DEVICE_INFO   = "/device"
	PATH_DEVICE_ADD    = utils.PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_INFO + utils.PATH_ADD
	PATH_DEVICE_DELETE = utils.PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_INFO + utils.PATH_DELETE
	PATH_DEVICE_UPDATE = utils.PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_INFO + utils.PATH_UPDATE
	PATH_DEVICE_QUERY  = utils.PATH_IOT_DEVICE_SINGLE + PATH_DEVICE_INFO + utils.PATH_QUERY
)

//基本信息添加
func DeviceAdd(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}

//基本信息添加
func DeviceDelete(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}

//基本信息添加
func DeviceUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}

//基本信息添加
func DeviceQuery(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}
