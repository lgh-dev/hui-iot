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

// @Summary 编辑系统用户

// @Description 编辑系统用户 账号、角色不予更新；只更新密码和姓名

// @Tags 系统用户模块

// @Param userName formData string true "账号"

// @Param password formData string false "密码"

// @Param realName formData string true "用户姓名"

// @Success 200 {string} json "{"msg": "添加成功"}"

// @Failure 401 {string} json "{"msg": "err.Error()"}"

// @Failure 402 {string} json "{"msg": "账号已存在"}"

// @Failure 403 {string} json "{"msg": "添加失败"}"

// @Router /user/update [post]
func DeviceTypeQuery(c *gin.Context) {
	c.JSON(http.StatusOK, domain.BuildSucc(&domain.ResultDTO{Data: domain.DeviceTypeMap}))
}
