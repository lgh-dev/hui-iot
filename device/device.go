package main

import (
	"github.com/gin-gonic/gin"
	. "light-iot/base/domain"
)

func main() {

	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, ResultDTO{Code: 1, Data: nil, Msg: "succ"})
	})
	r.Run(":9000")
}
