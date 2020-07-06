package listener

import (
	"github.com/panjf2000/ants/v2"
	"hui-iot/iot-worker/db"
	"hui-iot/iot-worker/utils"
	"strconv"
	"sync"
	"time"
)

//持久化事件监听
type PersistenceListener struct{}

var DataDB = *db.NewTDengine()

var pool *ants.PoolWithFunc

var once = sync.Once{}

func (pl *PersistenceListener) getPool() *ants.PoolWithFunc {
	once.Do(func() {
		if pool == nil {
			// Use the pool with a function,
			// set 10 to the capacity of goroutine pool and 1 second for expired duration.
			pool, _ = ants.NewPoolWithFunc(30, pl.ToSave)
		}
	})
	return pool
}

func (pl *PersistenceListener) ToSave(dto interface{}) {
	eventDTO := dto.(*EventDTO)
	var data = eventDTO.Event.(*map[string]float64)
	for uKey, value := range *data {
		save(eventDTO, &uKey, &value)
	}
}

func (pl *PersistenceListener) DoEvent(dto *EventDTO) {
	if &dto == nil || dto.EventType != EventTypeRead {
		return
	}
	//single insert
	pl.ToSave(dto)
	// Multithreading insert
	//pl.getPool().Invoke(dto)
}

//INSERT INTO hui_iot.read_0001_1000_temp
//USING hui_iot.iot_read_smart_car_camera
//TAGS ('smart_car_camera','0001_1000', 'temp')
//VALUES ('2018-01-04 00:00:00.000',95);
func save(dto *EventDTO, uKey *string, value *float64) {
	//TODO 先实现基于TDengine的,后续改成支持多种数据库
	sql := GenerateSQL(dto.DeviceModel, dto.GatewayUID, uKey, value)
	//log.Printf("insert read attr sql:%s", sql)
	err := DataDB.Insert(sql)
	utils.CheckErrorLn("插入失敗:%s", err)
}

func GenerateSQL(deviceModel *string, gatewayUID *string, readUkey *string, value *float64) *string {
	sql := "INSERT INTO hui_iot.read_" + *gatewayUID + "_" + *readUkey + " USING hui_iot.iot_read_" + *deviceModel + " TAGS ('" + *deviceModel + "','" + *gatewayUID + "', '" + *readUkey + "') VALUES ('" + time.Now().Format("2006-01-02 15:04:05.000000") + "'," + strconv.FormatFloat(*value, 'f', 9, 64) + ")\n"
	return &sql
}
