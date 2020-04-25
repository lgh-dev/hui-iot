package service

import (
	"github.com/jinzhu/gorm"
	"hui-iot/iot-server/config"
	. "hui-iot/iot-server/domain"
)

func GetGorm() *gorm.DB {
	db, err := gorm.Open("mysql", config.Conf.GetString("mysql.dbInfo"))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return db
}

func addDevice(device Device) {
	db := GetGorm()
	defer db.Close()
	deviceMgr := DeviceMgr(db)
	device.ID = IDTool.NextIdStr()
	deviceMgr.Create(&device)
}
