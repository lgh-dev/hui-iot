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

func DevicePUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, domain.BuildSucc(&domain.ResultDTO{}))
}

func DevicePQuery(c *gin.Context) {
	c.JSON(http.StatusOK, domain.BuildSucc(&domain.ResultDTO{}))
}
