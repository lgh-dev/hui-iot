package controller

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/gin-gonic/gin"
	"hui-iot/iot-server/common"
	"net/http"
)

//基本信息添加
func AddDevice(c *gin.Context) {
	c.JSON(http.StatusOK, common.BuildSucc(&common.ResultDTO{}))
}

//基本信息添加
func DeleteDevice(c *gin.Context) {
	c.JSON(http.StatusOK, common.BuildSucc(&common.ResultDTO{}))
}

//基本信息添加
func UpdateDevice(c *gin.Context) {
	c.JSON(http.StatusOK, common.BuildSucc(&common.ResultDTO{}))
}

//基本信息添加
func GetDevice(c *gin.Context) {
	c.JSON(http.StatusOK, common.BuildSucc(&common.ResultDTO{}))
}
