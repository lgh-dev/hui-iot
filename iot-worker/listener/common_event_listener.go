package listener

//事件类型
const (
	EventTypeRead            = "EVENT_TYPE_READ"
	EventTypeStatusOnline    = "EVENT_TYPE_STATUS_ONLINE"
	EventTypeStatusHeartbeat = "EVENT_TYPE_STATUS_HEARTBEAT"
	EventTypeStatusOffline   = "EVENT_TYPE_STATUS_OFFLINE"
	EventTypeCommandSend     = "EVENT_TYPE_COMMAND_SEND"
	EventTypeCommandReturn   = "EVENT_TYPE_COMMAND_RETURN"
	EventTypeAlarm           = "EVENT_TYPE_ALARM"
	EventTypeLog             = "EVENT_TYPE_LOG"
)

// 事件包装结体
type EventDTO struct {
	DeviceModel *string
	GatewayUID  *string
	EventType   string
	Event       interface{}
}

type Listener interface {
	DoEvent(EventDTO)
}
