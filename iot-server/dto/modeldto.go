package dto

import (
	"hui-iot/iot-server/domain"
	"time"
)

type AbsEntityDTO struct {
	ID         string    `json:"_id"`        // 主键 分布式唯一ID
	CreateBy   string    `json:"createBy"`   // 创建者
	CreateTime time.Time `json:"createTime"` //创建时间
	UpdateBy   string    `json:"updateBy"`   //更新者
	UpdateTime time.Time `json:"updateTime"` //更新时间
	IsDelete   bool      `json:"isDelete"`   // 软删除建,true为删除。
}

// DeviceModel
type DeviceModelDTO struct {
	ID               string             `yaml:"id"`             //唯一标识
	Version          string             `yaml:"version"`        //版本
	Auth             string             `yaml:"auth"`           //作者
	Name             string             `yaml:"name"`           //名称
	Desc             string             `yaml:"desc"`           //描述
	InvariantAttr    []AttrDefineDTO    `yaml:"invariant-attr"` //固定属性
	ReadAttrDefine   []AttrDefineDTO    `yaml:"read-attr"`      //只读属性
	ConfigAttrDefine []AttrDefineDTO    `yaml:"config-attr"`    //配置属性
	EventDefine      []EventDefineDTO   `yaml:"event"`          //事件定义
	CommandDefine    []CommandDefineDTO `yaml:"command"`        //指令定义
}

//属性定义
type AttrDefineDTO struct {
	UKey         string      `yaml:"uKey"`
	Name         string      `yaml:"name"`
	Required     bool        `yaml:"required"`
	DefaultValue interface{} `yaml:"defaultValue"`
	DataType     string      `yaml:"date-type"` // 字段类型：默认字符串
	Unit         string      `yaml:"unit"`      // 单位：默认值无
}

//指令定义
type CommandDefineDTO struct {
	UKey           string          `yaml:"uKey"`
	Name           string          `yaml:"name"`
	InParamsDefine []AttrDefineDTO `yaml:"in-params"` //配置属性
}

//事件定义
type EventDefineDTO struct {
	UKey            string          `yaml:"uKey"`
	Name            string          `yaml:"name"`
	Level           string          `yaml:"level"`
	OutParamsDefine []AttrDefineDTO `yaml:"out-params"`
}

// Device 物联网设备表;
type DeviceDTO struct {
	AbsEntityDTO
	Uid              string                 `json:"uid"`              // 设备UID
	Name             string                 `json:"name"`             // 设备名称
	GatewayID        string                 `json:"gatewayID"`        // 网关ID
	DeviceModelID    string                 `json:"deviceModelID"`    // 设备模型ID-外键 设备模型ID-外键
	OnlineStatus     string                 `json:"OnlineStatus"`     // 联网状态 在线ONLINE、离线OFFLINE、未激活NONACTIVE
	InvariantAttrMap map[string]KeyValueDTO `json:"invariantAttrMap"` //固定属性，不变的属性。
	ConfigAttrMap    map[string]KeyValueDTO `json:"configAttrMap"`    //配置属性，修改后动态更新到设备端的属性，如心跳间隔。
	SensorAttrMap    map[string]KeyValueDTO `json:"sensorAttrMap"`    //传感属性，动态上报的属性。
}

// DeviceCommand 指令下发表
type CommandDTO struct {
	AbsEntityDTO
	CommandID  string   `json:"commandID"`  // 指令ID
	DeviceID   string   `json:"deviceID"`   // 设备ID
	SendData   string   `json:"sendData"`   // 发送数据
	ReturnData string   `json:"returnData"` // 响应数据（同步或异步）
	Status     string   `json:"status"`     // 指令状态(send【下发】、ack【设备已接收】、succ【执行成功】、error【异常】)
	Msg        string   `json:"msg"`        // 执行信息，包括成功或失败的信息
	Retry      RetryDTO `json:"retry"`      //重试机制，就不重试
}

//属性标识和值
type KeyValueDTO struct {
	UKey  string `json:"uKey"`  // 属性标识
	Value string `json:"value"` // 属性值

}

// DeviceEvent 设备事件表
type EventDTO struct {
	AbsEntityDTO
	Time        int64          `json:"time"`        // 落库时间 时序数据库的主键
	SendTime    time.Time      `json:"sendTime"`    // 事件发送时间
	EventDefine EventDefineDTO `json:"eventDefine"` // 事件定义标识
	Info        string         `json:"info"`        // 事件信息
	Status      string         `json:"status"`      // 事件状态 事件状态（告警ALARM，结束OVER）
}

//重试对象
type RetryDTO struct {
	Interval        string `json:"interval"`        // 重发间隔(s)
	Number          string `json:"number"`          // 重发次数
	CompletedNumber string `json:"completedNumber"` // 已重发次数
}

type ApplyDTO struct {
	uKey    string  `yaml:"uKey"`    //唯一标识
	Version string  `yaml:"version"` //版本
	Auth    AuthDTO `yaml:"auth"`    //权限信息
	author  string  `yaml:"author"`  //作者
	Desc    string  `yaml:"desc"`    //描述
}

type AuthDTO struct {
	AppKey string `yaml:"appKey"` // 应用账号
	Secret string `yaml:"secret"` // 应用密钥
}

func (deviceDTO *DeviceDTO) ToDevice() domain.Device {
	device := domain.Device{}
	device.ID = deviceDTO.ID
	device.Name = deviceDTO.Name
	//@TODO 待完成
	return device
}
