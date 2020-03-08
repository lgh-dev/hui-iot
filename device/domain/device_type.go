package domain

import (
	"github.com/edwingeng/wuid/callback/wuid"
	"gopkg.in/yaml.v2"
	. "hui-iot/base/utils"
	"io/ioutil"
	"time"
)

// config file path
const configPath = "../config/devicetype/"

//The config file prefix of key
type PrefixKey string

// save device type info
var DeviceTypeMap = make(map[string]DeviceType)

// base entity
type BaseEntity struct {
	ID         int64 //系统唯一标识
	CreateTime time.Time
	UpdateTime time.Time
}

var idUtil *wuid.WUID

//动态生成表和对应的sql
type DynamicSQL struct {
	TableName      string //表名 ，格式：iot_[DeviceTypeID]_[PAttr|RAttr|CAttr|Func]
	CreateTableSQL string //
	DropTableSQL   string
	InsertSQL      string
	QuerySQL       string
	UpdateSQL      string
	DeleteSQL      string
}

// DeviceType
type DeviceType struct {
	ID               string        `yaml:"ID"`            //唯一标识
	InvariantDefine  []AttrDefine  `yaml:"invarant-attr"` //固定属性
	ReadDefine       []AttrDefine  `yaml:"read-attr"`     //只读属性
	ConfigDefine     []AttrDefine  `yaml:"config-attr"`   //配置属性
	EventDefine      []EventDefine `yaml:"event"`         //功能函数
	Function         []Function    `yaml:"function"`      //功能函数
	InvariantAttrSQL DynamicSQL    //固定属性动态sql
	ReadAttrSQL      DynamicSQL    //动态sql
	ConfigAttrSQL    DynamicSQL
	FuncSQL          DynamicSQL
}

//属性定义
type AttrDefine struct {
	Key          string      `yaml:"key"`
	Name         string      `yaml:"name"`
	Required     bool        `yaml:"required"`
	DefaultValue interface{} `yaml:"defaultValue"`
	DataType     string      `yaml:"date-type"` // 字段类型：默认字符串
	Unit         string      `yaml:"unit"`      // 单位：默认值无
}

type Function struct {
	Key            string       `yaml:"key"`
	Name           string       `yaml:"name"`
	InParamsDefine []AttrDefine `yaml:"in-params"` //配置属性
}

type EventDefine struct {
	Key             string       `yaml:"key"`
	Name            string       `yaml:"name"`
	Level           string       `yaml:"level"`
	OutParamsDefine []AttrDefine `yaml:"out-params"`
}

type ResultDTO struct {
	Code int         `json:"code"` //Succ:1  or err code :<1
	Data interface{} `json:"data"` //Result data
	Msg  string      `json:"msg"`  //Succ or Err msg
}

func BuildSucc(result *ResultDTO) *ResultDTO {
	result.Code = 1
	if result.Msg == "" {
		result.Msg = "succ"
	}
	return result
}

func BuildError(result *ResultDTO) *ResultDTO {
	result.Code = 0
	if result.Msg == "" {
		result.Msg = "error"
	}
	return result
}

//初始化配置文件
func init() {
	GetDeviceTypes() //Init device type config info.
}

func GetDeviceTypes() {
	//读取配置文件目录。
	yamlFiles, err := ioutil.ReadDir(configPath)
	if err != nil {
		Log.Error("Read config path ${root}/base/config/ err:{}", err.Error())
	}
	//cycle read config yaml files transform device type info
	for i := 0; i < len(yamlFiles); i++ {
		file := yamlFiles[i]
		if !file.IsDir() {
			yamlFile, err := ioutil.ReadFile(configPath + file.Name())
			if err != nil {
				Log.Error("read yaml file err:{}", err.Error())
			}
			deviceType := initDeviceType(file.Name(), yamlFile)
			DeviceTypeMap[deviceType.ID] = deviceType
		}
	}
}

//初始化设备类型
func initDeviceType(filename string, bytes []byte) DeviceType {
	Log.Debug("init yaml file name is ", filename)
	deviceType := DeviceType{}
	err := yaml.Unmarshal(bytes, &deviceType)
	if err != nil {
		panic("Read config file of device-type err: " + filename + "" + err.Error())
	}
	//Is err
	return deviceType
}
