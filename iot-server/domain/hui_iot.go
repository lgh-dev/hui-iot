package domain

import (
	"time"
)

// Device 物联网设备表;
type Device struct {
	ID            string `gorm:"primary_key;column:id;type:varchar(32);not null" json:"-"`                // 主键 分布式唯一ID
	DeviceUId     string `gorm:"column:device_uid;type:varchar(32);not null" json:"device_uid"`           // 设备UID
	DeviceName    string `gorm:"column:device_name;type:varchar(1024)" json:"device_name"`                // 设备名称
	GatewayID     string `gorm:"column:gateway_id;type:varchar(32)" json:"gateway_id"`                    // 网关ID
	DeviceModelID string `gorm:"column:device_model_id;type:varchar(32);not null" json:"device_model_id"` // 设备模型ID-外键 设备模型ID-外键
	OnlineStatus  int    `gorm:"column:online_status;type:int;not null" json:"online_status"`             // 联网状态 在线（1）、离线（0）、未激活（-1）
}

// DeviceCommand 指令下发表
type DeviceCommand struct {
	CmdID                string `gorm:"column:cmd_id;type:varchar(32)" json:"cmd_id"`                                 // 指令ID
	DeviceID             string `gorm:"column:device_id;type:varchar(32)" json:"device_id"`                           // 设备ID
	SendData             string `gorm:"column:send_data;type:varchar(32)" json:"send_data"`                           // 发送数据
	ReturnData           string `gorm:"column:return_data;type:varchar(32)" json:"return_data"`                       // 响应数据（同步或异步）
	Status               string `gorm:"column:status;type:varchar(32)" json:"status"`                                 // 指令状态(send【下发】、ack【设备已接收】、succ【执行成功】、error【异常】)
	Msg                  string `gorm:"column:msg;type:varchar(32)" json:"msg"`                                       // 执行信息，包括成功或失败的信息
	IsRetry              string `gorm:"column:is_retry;type:varchar(1)" json:"is_retry"`                              // 是否需要重发
	RetryInterval        string `gorm:"column:retry_interval;type:varchar(32)" json:"retry_interval"`                 // 重发间隔(s)
	RetryNumber          string `gorm:"column:retry_number;type:varchar(32)" json:"retry_number"`                     // 重发次数
	RetryCompletedNumber string `gorm:"column:retry_completed_number;type:varchar(32)" json:"retry_completed_number"` // 已重发次数
}

// DeviceConfigAttr 设备配置属性表
type DeviceConfigAttr struct {
	ID       int64  `gorm:"primary_key;column:id;type:bigint;not null" json:"-"`    // ID 主键
	DeviceID int64  `gorm:"column:device_id;type:bigint;not null" json:"device_id"` // 设备ID-外键
	Ukey     string `gorm:"column:ukey;type:varchar(32);not null" json:"ukey"`      // 配置属性标识
	Value    string `gorm:"column:value;type:varchar(3072)" json:"value"`           // 配置属性-值
}

// DeviceEvent 设备事件表
type DeviceEvent struct {
	Time     int64     `gorm:"primary_key;column:time;type:bigint;not null" json:"time"` // 落库时间 时序数据库的主键
	SendTime time.Time `gorm:"column:send_time;type:datetime;not null" json:"send_time"` // 事件发送时间
	Ukey     int64     `gorm:"column:ukey;type:bigint;not null" json:"ukey"`             // 事件定义标识
	Info     string    `gorm:"column:info;type:text;not null" json:"info"`               // 事件信息
	Level    string    `gorm:"column:level;type:varchar(32);not null" json:"level"`      // 事件级别(信息INFO、告警ALARM、故障ERROR)
	Status   string    `gorm:"column:status;type:varchar(32);not null" json:"status"`    // 事件状态 事件状态（告警ALARM，结束OVER）
}

// DeviceInvariantAttr 设备固定属性表
type DeviceInvariantAttr struct {
	ID       int64  `gorm:"primary_key;column:id;type:bigint;not null" json:"-"`    // ID 主键
	DeviceID int64  `gorm:"column:device_id;type:bigint;not null" json:"device_id"` // 设备ID-外键 设备ID-外键
	Ukey     string `gorm:"column:ukey;type:varchar(32);not null" json:"ukey"`      // 固定属性标识
	Value    string `gorm:"column:value;type:varchar(3072)" json:"value"`           // 固定属性-值
}
