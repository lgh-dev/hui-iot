package listener

import (
	"hui-iot/iot-worker/utils"
	"testing"
)

var listener = PersistenceListener{}

func BenchmarkGenerateSQL(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			deviceModel := "smart_car_camera"
			gatewayUID := "1100101"
			uKey := "temp"
			value := float64(123)
			GenerateSQL(&deviceModel, &gatewayUID, &uKey, &value)
		}
	})
}

func TestPersistenceListener_DoEvent(t *testing.T) {

	//for i := 0; i < 100000; i++ {
	//	dataMap:=make(map[string]float64)
	//	dataMap["temp"]=float64(123)
	//	//dataMap["msgId"]=float64(223)
	//	deviceModel:="smart_car_camera"
	//	gatewayUID:="1100101"
	//	eventDTO:=EventDTO{DeviceModel:&deviceModel,GatewayUID:&gatewayUID,EventType: EventTypeRead,Event:&dataMap}
	//	listener.DoEvent(&eventDTO)
	//}
	defer DataDB.Connect.Close()
	for i := 0; i < 100000; i++ {
		dataMap := make(map[string]float64)
		dataMap["temp"] = float64(123)
		//dataMap["msgId"]=float64(223)
		deviceModel := "smart_car_camera"
		gatewayUID := "1100101"
		eventDTO := EventDTO{DeviceModel: &deviceModel, GatewayUID: &gatewayUID, EventType: EventTypeRead, Event: &dataMap}
		var data = eventDTO.Event.(*map[string]float64)
		for uKey, value := range *data {
			//TODO 先实现基于TDengine的,后续改成支持多种数据库
			sql := GenerateSQL(eventDTO.DeviceModel, eventDTO.GatewayUID, &uKey, &value)
			//log.Printf("insert read attr sql:%s", sql)
			err := DataDB.Insert(sql)
			utils.CheckErrorLn("插入失敗:%s", err)
		}

	}
}

func BenchmarkPersistenceListener_DoEvent(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.Run("doevent", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dataMap := make(map[string]float64)
			dataMap["temp"] = float64(123)
			//dataMap["msgId"]=float64(223)
			deviceModel := "smart_car_camera"
			gatewayUID := "1100101"
			eventDTO := EventDTO{DeviceModel: &deviceModel, GatewayUID: &gatewayUID, EventType: EventTypeRead, Event: &dataMap}
			listener.DoEvent(&eventDTO)
		}
	})
	//b.RunParallel(func(pb *testing.PB) {
	//	for pb.Next(){
	//		dataMap:=make(map[string]float64)
	//		dataMap["temp"]=float64(123)
	//		//dataMap["msgId"]=float64(223)
	//		deviceModel:="smart_car_camera"
	//		gatewayUID:="1100101"
	//		eventDTO:=EventDTO{DeviceModel:&deviceModel,GatewayUID:&gatewayUID,EventType: EventTypeRead,Event:&dataMap}
	//		listener.DoEvent(&eventDTO)
	//	}
	//})
}
