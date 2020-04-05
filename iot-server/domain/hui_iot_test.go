package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"hui-iot/iot-server/config"
	"testing"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

/**
数据库连接信息
*/

func GetGorm(dataSourceName string) *gorm.DB {
	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return db
}

// TestFuncGet 测试条件获(Get/Gets)
func TestFuncGetAndCreateAndDelete(t *testing.T) {
	OpenRelated() // 打开全局预加载 (外键)
	conf := config.GetConf()
	db := GetGorm(conf.GetString("mysql.dbInfo"))
	defer db.Close()

	iotDevice := IotDevice{ID: 1, DeviceUId: "001", DeviceName: "test001", DeviceModelID: 1}
	iotDeviceMgr := IotDeviceMgr(db.Where("id = ?", 1))
	iotDeviceMgr.SingularTable(true)    //表名不加复数。
	iotDeviceMgr.Create(&iotDevice)     // 插入新的。
	iotDevice2, _ := iotDeviceMgr.Get() // 单条获取
	fmt.Println(iotDevice2)
	assert.Equal(t, "test001", iotDevice2.DeviceName)
	iotDeviceMgr.Delete(&iotDevice)     //删除
	iotDevice3, _ := iotDeviceMgr.Get() // 单条获取
	assert.NotEqual(t, "test001", iotDevice3.DeviceName)

}
