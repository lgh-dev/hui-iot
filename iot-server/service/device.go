package service

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/gin-gonic/gin"
	. "hui-iot/iot-server/domain"
	"sync"
)

var once sync.Once

//基本信息添加
func AddDevice(device Device) error {

	return nil
}

//基本信息添加
func DeleteDevice(ids []string) error {

	return nil

}

//基本信息添加
func UpdateDevice() {

}

//基本信息添加
func FindDeviceById(c *gin.Context) {

}

//基本信息添加
func FindDeviceByPage(c *gin.Context) {

}
