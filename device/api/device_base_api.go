package api

import (
	"github.com/gin-gonic/gin"
	. "light-iot/base/domain"
	"net/http"
)

func deviceBaseAdd(c *gin.Context) {
	c.JSON(http.StatusOK, BuildSucc(&ResultDTO{}))
}
