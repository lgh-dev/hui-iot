package middleware

import (
	"github.com/gin-gonic/gin"
	"hui-iot/iot-server/config"
	"hui-iot/iot-server/service"
	"log"
	"strings"
	"time"
)

func AppAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		accessTokenAll := c.GetHeader("Authorization")

		if accessTokenAll != "" && strings.HasPrefix(accessTokenAll, "Bearer") {
			accessToken := accessTokenAll[7:]
			log.Printf("accessToken succ:%s\n", accessToken)
			claims, err := service.ParseToken(accessToken)
			if err != nil || claims == nil {
				log.Printf("accessToken parse err %s", accessToken)
				c.JSON(401, "accessToken is illegal！")
				c.Abort()
				return
			}
			//accessToken时间未过期且启用了该应用。
			if claims.ExpiresAt > time.Now().Unix() {
				for _, apple := range config.AppYaml.Apples {
					if apple.AppKey == claims.AppKey {
						c.Header("appKey", claims.AppKey)
					}
				}
			}
		}
	}
}
