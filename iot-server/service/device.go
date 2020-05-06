package service

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/globalsign/mgo/bson"
	. "hui-iot/iot-server/dao"
	. "hui-iot/iot-server/domain"
	"hui-iot/iot-server/dto"
	"time"
)

type DeviceService struct {
	ICurdService
}

func (deviceService *DeviceService) Add(Entity interface{}) (string, error) {
	device := Entity.(dto.DeviceDTO)
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
	return device.ID, c.Insert(&device)
}
func (deviceService *DeviceService) FindByID(id string) interface{} {
	session := CloneSession()
	defer session.Close()
	c := getCollection(session)
	return c.Find(bson.M{"_id": id, "isDelete": false}).One(&Device{})
}
func (deviceService *DeviceService) FindByIDs(ids []string) interface{} {
	session := CloneSession()
	defer session.Close()
	c := getCollection(session)
	return c.Find(bson.M{"_id": bson.M{"$in": ids}, "isDelete": false}).All(&Device{})
}
func (deviceService *DeviceService) FindByPage(condition interface{}, page Page) interface{} {
	session := CloneSession()
	defer session.Close()
	c := getCollection(session)
	//TODO 查询条件
	return c.Find(bson.M{}).Skip((page.PageNumber - 1) * page.PageSize).Limit(page.PageSize).All(&Device{})

}
func (deviceService *DeviceService) Update(Entity interface{}) error {
	device := Entity.(Device)
	session := CloneSession()
	defer session.Close()
	c := getCollection(session)
	_, err := c.UpsertId(device.ID, device)
	return err
}
func (deviceService *DeviceService) Delete(ids []string) error {
	session := CloneSession()
	defer session.Close()
	c := getCollection(session)
	return c.Remove(bson.M{"_id": bson.M{"$in": ids}})
}

func (deviceService *DeviceService) ExistsDeviceByName(name string) bool {
	session := CloneSession()
	defer session.Close()
	c := getCollection(session)
	var device Device
	c.Find(bson.D{{"name", name}, {"isDelete", false}}).One(&device)
	if device.Name == "" {
		return false
	}
	return true
}
