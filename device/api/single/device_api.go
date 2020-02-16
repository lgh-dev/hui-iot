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
func AddDevice(c *gin.Context) {
	device := domain.Device{}
	c.BindJSON(device)
	c.JSON(http.StatusOK, domain.BuildSucc(&domain.ResultDTO{}))
}

//基本信息添加
func DeleteDevice(c *gin.Context) {
	c.JSON(http.StatusOK, domain.BuildSucc(&domain.ResultDTO{}))
}

//基本信息添加
func UpdateDevice(c *gin.Context) {
	c.JSON(http.StatusOK, domain.BuildSucc(&domain.ResultDTO{}))
}

//基本信息添加
func GetDevice(c *gin.Context) {
	c.JSON(http.StatusOK, domain.BuildSucc(&domain.ResultDTO{}))
}
