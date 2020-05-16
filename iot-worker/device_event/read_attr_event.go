package device_event

import (
	"encoding/json"
	"github.com/thinkeridea/go-extend/exstrings"
	"hui-iot/iot-worker/db"
	"hui-iot/iot-worker/iotevent"
	"hui-iot/iot-worker/utils"
	"log"
	"strconv"
	"time"
)

//请求Topic：/hiot/sys/{deviceModelID}/{deviceID}/dv/r/up
//响应Topic：/hiot/sys/{deviceModelID}/{deviceID}/dv/r/up_reponse
var ReadAttrEventOperation = &iotevent.Operation{Context: &iotevent.Context{
	FromContext: iotevent.FromContext{
		FromTopic: "/hiot/sys/+/+/dv/r/up",
		FromQos:   0,
	},
}, EventHandler: ReadAttrEventHandler}

var DataDB = db.NewTDengineRestful()
var num = 0

//topic:/hiot/sys/car/1111/dv/r/up
//payload:{"temp":34}
func ReadAttrEventHandler(c *iotevent.Context) *iotevent.Context {
	num++
	if num%100 == 0 {
		log.Printf("累计消息数量：%d", num)
	}
	deviceModel := utils.GetStrForTopic(c.Topic, 3)
	gatewayUID := utils.GetStrForTopic(c.Topic, 4)
	data := make(map[string]float64)
	err := json.Unmarshal(c.Payload, &data)
	if err != nil {
		log.Printf("解析json错误，只读数据格式：%s", c.Payload)
	}
	for uKey, value := range data {
		sql := GenerateSQL(deviceModel, gatewayUID, uKey, value)
		insert(sql)
	}
	return c
}

//INSERT INTO hui_iot.read_0001_1000_temp
//USING hui_iot.iot_read_smart_car_camera
//TAGS ('smart_car_camera','0001_1000', 'temp')
//VALUES ('2018-01-04 00:00:00.000',95);
func insert(sql *string) {
	//log.Printf("insert read attr sql:%s", *sql)
	//if err := DataDB.Insert("INSERT INTO " +*sql); err != nil {
	//	log.Printf("插入失敗:%s", err)
	//}
}

var chanSQL = make(chan string, 10)

func GenerateSQL(deviceModel string, gatewayUID string, readUkey string, value float64) *string {
	sql := "hui_iot.read_{gatewayUID}_{readUkey} USING hui_iot.iot_read_{deviceModel} TAGS ('{deviceModel}','{gatewayUID}', '{readUkey}') VALUES ('{timestamp}',{value})\n"
	sql = exstrings.Replace(sql, "{gatewayUID}", gatewayUID, 2)
	sql = exstrings.Replace(sql, "{deviceModel}", deviceModel, 2)
	sql = exstrings.Replace(sql, "{readUkey}", readUkey, 2)
	sql = exstrings.Replace(sql, "{timestamp}", time.Now().Format("2006-01-02 15:04:05.000"), 1)
	sql = exstrings.Replace(sql, "{value}", strconv.FormatFloat(value, 'f', 9, 64), 1)
	return &sql
}
func BatchInsert(sql string) {
	insertSQL := "INSERT INTO "
	go func(sql string) {
		chanSQL <- sql
	}(sql)

	go func() {
		var batchSQL string
		for i := 0; i < 10; i++ {
			batchSQL += <-chanSQL
		}
		batchSQL = insertSQL + batchSQL + ";"
		insert(&batchSQL)
	}()

}
