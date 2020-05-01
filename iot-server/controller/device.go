package controller

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/gin-gonic/gin"
	"hui-iot/iot-server/dto"
	"hui-iot/iot-server/service"
	"net/http"
	"sync"
)

var once sync.Once

//基本信息添加
func AddDevice(c *gin.Context) {
	var deviceDTO dto.DeviceDTO
	if err := c.ShouldBind(&deviceDTO); err != nil {
		c.JSON(http.StatusOK, dto.BuildSucc(&dto.ResultDTO{Data: err.Error()}))
		c.Abort()
		return
	}
	service.AddDevice(deviceDTO.ToDevice())
	c.JSON(http.StatusOK, dto.BuildSucc(&dto.ResultDTO{}))
}

//基本信息添加
func DeleteDevice(c *gin.Context) {
	id := c.Param("id")
	service.DeleteDevice([]string{id})
	c.JSON(http.StatusOK, dto.BuildSucc(&dto.ResultDTO{}))
}

//基本信息添加
func UpdateDevice(c *gin.Context) {
	deviceDTO := dto.DeviceDTO{}
	c.BindJSON(&deviceDTO)
	service.UpdateDevice(deviceDTO.ToDevice())
	c.JSON(http.StatusOK, dto.BuildSucc(&dto.ResultDTO{}))
}

//基本信息添加
func FindDeviceById(c *gin.Context) {
	id := c.Param("id")
	deviceDTO := service.FindDeviceById(id)
	c.JSON(http.StatusOK, dto.BuildSucc(&dto.ResultDTO{Data: &deviceDTO}))
}

//基本信息添加
func FindDeviceByPage(c *gin.Context) {
	c.JSON(http.StatusOK, dto.BuildSucc(&dto.ResultDTO{}))
}

//基本信息添加
func SendCommand(c *gin.Context) {
	c.JSON(http.StatusOK, dto.BuildSucc(&dto.ResultDTO{}))
}
