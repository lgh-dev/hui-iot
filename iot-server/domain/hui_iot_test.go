package domain

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"hui-iot/iot-server/config"
	"testing"
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
	db := GetGorm(config.Conf.GetString("mysql.dbInfo"))
	defer db.Close()

	iotDevice := Device{ID: "1", DeviceUId: "001", DeviceName: "test001", DeviceModelID: "1"}
	deviceMgr := DeviceMgr(db.Where("id = ?", 1))
	deviceMgr.SingularTable(true)    //表名不加复数。
	deviceMgr.Create(&iotDevice)     // 插入新的。
	iotDevice2, _ := deviceMgr.Get() // 单条获取
	fmt.Println(iotDevice2)
	assert.Equal(t, "test001", iotDevice2.DeviceName)
	deviceMgr.Delete(&iotDevice)     //删除
	iotDevice3, _ := deviceMgr.Get() // 单条获取
	assert.NotEqual(t, "test001", iotDevice3.DeviceName)

}
