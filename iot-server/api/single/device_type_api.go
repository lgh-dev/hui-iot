package single

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/gin-gonic/gin"
)

// @description 通过id获取用户
// @version 1.0
// @accept application/x-www-form-urlencoded
// @param id path int true "id"
// @router /api/svc-iot-iot-server/v1/devicetype/:id [get]
//func DeviceTypeQuery(c *gin.Context) {
//	c.JSON(http.StatusOK, common.BuildSucc(&common.ResultDTO{Data: common.DeviceTypeMap}))
//}

// @description 通过id获取用户
// @version 1.0
// @accept application/x-json-stream
// @param id path int true "id"
// @router /user/{id} [get]
func UserGet(ctx *gin.Context) {
	//...
}
