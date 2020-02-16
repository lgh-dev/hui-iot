package single

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/gin-gonic/gin"
	"hui-iot/device/domain"
	"net/http"
)

//基本信息添加
func DeviceAdd(c *gin.Context) {
	c.JSON(http.StatusOK, domain.BuildSucc(&domain.ResultDTO{}))
}

//基本信息添加
func DeviceDelete(c *gin.Context) {
	c.JSON(http.StatusOK, domain.BuildSucc(&domain.ResultDTO{}))
}

//基本信息添加
func DeviceUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, domain.BuildSucc(&domain.ResultDTO{}))
}

//基本信息添加
func DeviceQuery(c *gin.Context) {
	c.JSON(http.StatusOK, domain.BuildSucc(&domain.ResultDTO{}))
}
