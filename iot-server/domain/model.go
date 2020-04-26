package domain

import (
	"time"
)

type AbsEntity struct {
	ID         string    `json:"_id"`        // 主键 分布式唯一ID
	CreateBy   string    `json:"createBy"`   // 创建者
	CreateTime time.Time `json:"createTime"` //创建时间
	UpdateBy   string    `json:"updateBy"`   //更新者
	UpdateTime time.Time `json:"updateTime"` //更新时间
	IsDelete   bool      `json:"isDelete"`   // 软删除建,true为删除。
}

// DeviceModel
type DeviceModel struct {
	ID               string          `yaml:"id"`             //唯一标识
	Version          string          `yaml:"version"`        //版本
	Auth             string          `yaml:"auth"`           //作者
	Name             string          `yaml:"name"`           //名称
	Desc             string          `yaml:"desc"`           //描述
	InvariantAttr    []AttrDefine    `yaml:"invariant-attr"` //固定属性
	ReadAttrDefine   []AttrDefine    `yaml:"read-attr"`      //只读属性
	ConfigAttrDefine []AttrDefine    `yaml:"config-attr"`    //配置属性
	EventDefine      []EventDefine   `yaml:"event"`          //事件定义
	CommandDefine    []CommandDefine `yaml:"command"`        //指令定义
}

//属性定义
type AttrDefine struct {
	UKey         string      `yaml:"uKey"`
	Name         string      `yaml:"name"`
	Required     bool        `yaml:"required"`
	DefaultValue interface{} `yaml:"defaultValue"`
	DataType     string      `yaml:"date-type"` // 字段类型：默认字符串
	Unit         string      `yaml:"unit"`      // 单位：默认值无
}

//指令定义
type CommandDefine struct {
	UKey           string       `yaml:"uKey"`
	Name           string       `yaml:"name"`
	InParamsDefine []AttrDefine `yaml:"in-params"` //配置属性
}

//事件定义
type EventDefine struct {
	UKey            string       `yaml:"uKey"`
	Name            string       `yaml:"name"`
	Level           string       `yaml:"level"`
	OutParamsDefine []AttrDefine `yaml:"out-params"`
}

// Device 物联网设备表;
type Device struct {
	AbsEntity
	Uid              string              `json:"uid"`              // 设备UID
	Name             string              `json:"name"`             // 设备名称
	GatewayID        string              `json:"gatewayID"`        // 网关ID
	DeviceModelID    string              `json:"deviceModelID"`    // 设备模型ID-外键 设备模型ID-外键
	OnlineStatus     int                 `json:"OnlineStatus"`     // 联网状态 在线（1）、离线（0）、未激活（-1）
	InvariantAttrMap map[string]KeyValue `json:"invariantAttrMap"` //固定属性，不变的属性。
	ConfigAttrMap    map[string]KeyValue `json:"configAttrMap"`    //配置属性，修改后动态更新到设备端的属性，如心跳间隔。
	SensorAttrMap    map[string]KeyValue `json:"sensorAttrMap"`    //传感属性，动态上报的属性。
}

// DeviceCommand 指令下发表
type Command struct {
	AbsEntity
	CommandID  string `json:"commandID"`  // 指令ID
	DeviceID   string `json:"deviceID"`   // 设备ID
	SendData   string `json:"sendData"`   // 发送数据
	ReturnData string `json:"returnData"` // 响应数据（同步或异步）
	Status     string `json:"status"`     // 指令状态(send【下发】、ack【设备已接收】、succ【执行成功】、error【异常】)
	Msg        string `json:"msg"`        // 执行信息，包括成功或失败的信息
	Retry      Retry  `json:"retry"`      //重试机制，就不重试
}

//重试对象
type Retry struct {
	Interval        string `json:"interval"`        // 重发间隔(s)
	Number          string `json:"number"`          // 重发次数
	CompletedNumber string `json:"completedNumber"` // 已重发次数
}

//属性标识和值
type KeyValue struct {
	UKey  string `json:"uKey"`  // 属性标识
	Value string `json:"value"` // 属性值

}

// DeviceEvent 设备事件表
type Event struct {
	AbsEntity
	Time        int64       `json:"time"`        // 落库时间 时序数据库的主键
	SendTime    time.Time   `json:"sendTime"`    // 事件发送时间
	EventDefine EventDefine `json:"eventDefine"` // 事件定义标识
	Info        string      `json:"info"`        // 事件信息
	Status      string      `json:"status"`      // 事件状态 事件状态（告警ALARM，结束OVER）
}

type Apply struct {
	uKey    string `yaml:"uKey"`    //唯一标识
	Version string `yaml:"version"` //版本
	Auth    Auth   `yaml:"auth"`    //权限信息
	author  string `yaml:"author"`  //作者
	Desc    string `yaml:"desc"`    //描述
}

type Auth struct {
	AppKey string `yaml:"appKey"` // 应用账号
	Secret string `yaml:"secret"` // 应用密钥
}
