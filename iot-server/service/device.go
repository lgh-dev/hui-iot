package service

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/gin-gonic/gin"
	. "hui-iot/iot-server/dao"
	. "hui-iot/iot-server/domain"
	"time"
)

//基本信息添加
func AddDevice(device Device) error {
	session := CloneSession()
	defer session.Close()
	c := getCollection(session)
	device.ID = IDTool.NextIdStr()
	device.CreateBy = ""
	device.CreateTime = time.Now()
	device.UpdateBy = ""
	device.UpdateTime = time.Now()
	device.IsDelete = false
	device.OnlineStatus = NONACTIVE
	return c.Insert(&device)
}

//基本信息添加
func DeleteDevice(ids []string) error {
	return nil

}

//基本信息添加
func UpdateDevice() {

}

//基本信息添加
func FindDeviceById(id string) Device {
	session := CloneSession()
	defer session.Close()
	c := getCollection(session)
	device := Device{}
	c.FindId(id).One(&device)
	return device
}

//基本信息添加
func FindDeviceByPage(c *gin.Context) {

}
