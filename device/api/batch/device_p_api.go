package single

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/gin-gonic/gin"
	domain2 "hui-iot/base/common"
	"net/http"
)

func DevicePUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, domain2.BuildSucc(&domain2.ResultDTO{}))
}

func DevicePQuery(c *gin.Context) {
	c.JSON(http.StatusOK, domain2.BuildSucc(&domain2.ResultDTO{}))
}
