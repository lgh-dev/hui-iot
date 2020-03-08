package common

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// save device types info
var DeviceTypeMap = make(map[string]DeviceType)

//动态生成表和对应的sql
type DynamicSQL struct {
	TableName      string //表名 ，格式：iot_[DeviceTypeID]_[PAttr|RAttr|CAttr|Func]
	dbType         string //数据库类型 ，格式：[mysql,TDengine]
	CreateTableSQL string // 建表语句
	DropTableSQL   string // 删表语句
	InsertSQL      string // 插入语句
	QuerySQL       string // 查询语句
	UpdateSQL      string // 更新语句
	DeleteSQL      string // 删除语句
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

//功能函数
type Function struct {
	Key            string       `yaml:"key"`
	Name           string       `yaml:"name"`
	InParamsDefine []AttrDefine `yaml:"in-params"` //配置属性
}

//事件定义
type EventDefine struct {
	Key             string       `yaml:"key"`
	Name            string       `yaml:"name"`
	Level           string       `yaml:"level"`
	OutParamsDefine []AttrDefine `yaml:"out-params"`
}

func GetDeviceTypes(configPath string) bool {
	//读取配置文件目录。
	yamlFiles, err := ioutil.ReadDir(configPath)
	if err != nil {
		Log.Error("Read config path ${root}/base/config/ err:{}", err.Error())
		return false
	}
	//cycle read config yaml files transform device type info
	for i := 0; i < len(yamlFiles); i++ {
		file := yamlFiles[i]
		if !file.IsDir() {
			yamlFile, err := ioutil.ReadFile(configPath + file.Name())
			if err != nil {
				Log.Error("read yaml file err:{}", err.Error())
				return false
			}
			deviceType := getDeviceType(file.Name(), yamlFile)
			DeviceTypeMap[deviceType.ID] = deviceType
		}
	}
	return true
}

//初始化设备类型
func getDeviceType(filename string, bytes []byte) DeviceType {
	deviceType := DeviceType{}
	err := yaml.Unmarshal(bytes, &deviceType)
	if err != nil {
		panic("Read config file of device-type err: " + filename + "" + err.Error())
	}
	//初始化SQL
	readDynamicSQL := DynamicSQL{}
	readDynamicSQL.TableName = "read_attr_" + deviceType.ID
	readDynamicSQL.CreateTableSQL = ""
	//TODO 拼sql
	return deviceType
}
