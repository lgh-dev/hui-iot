package controller

import (
	"github.com/gin-gonic/gin"
	"hui-iot/iot-server/config"
	"hui-iot/iot-server/dto"
	"hui-iot/iot-server/service"
)

type ApplicationApi struct {
}

//获取Token
func (applicationApi *ApplicationApi) GetToken(c *gin.Context) {
	appKey := c.Query("appKey")
	secret := c.Query("secret")
	isTrue := false
	for _, apple := range config.AppYaml.Apples {
		if apple.AppKey == appKey && apple.Secret == secret {
			isTrue = true
		}
	}
	if isTrue {
		jwt, _ := service.GenerateToken(appKey)
		c.JSON(200, dto.BuildSucc(&dto.ResultDTO{Data: jwt}))
	}
	c.JSON(401, dto.BuildSucc(&dto.ResultDTO{Msg: "No permission. "}))

}
