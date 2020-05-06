package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func AppAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		acctokenAll := c.GetHeader("Authorization")

		isTrue := false

		if acctokenAll != "" && strings.HasPrefix(acctokenAll, "Bearer") {
			acctoken := acctokenAll[7:]
			isTrue = true
			log.Printf("accessToken succ:%s\n", acctoken)
			//TODO oauth2.0
		}
		if !isTrue {
			c.JSON(401, "accessToken is illegal！")
			c.Abort()
		}
	}

}
