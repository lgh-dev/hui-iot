package common

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

// save iot-server types info
var DeviceModelMap = make(map[string]DeviceModel)

var once = sync.Once{}

//动态生成表和对应的sql
type DynamicSQL struct {
	TableName      string //表名 ，格式：iot_[DeviceModelID]_[PAttr|RAttr|CAttr|Func]
	dbType         string //数据库类型 ，格式：[mysql,TDengine]
	CreateTableSQL string // 建表语句
	DropTableSQL   string // 删表语句
	InsertSQL      string // 插入语句
	QuerySQL       string // 查询语句
	UpdateSQL      string // 更新语句
	DeleteSQL      string // 删除语句
}

// DeviceModel
type DeviceModel struct {
	ID               string        `yaml:"ID"`            //唯一标识
	Version          string        `yaml:"version"`       //版本
	Auth             string        `yaml:"auth"`          //作者
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

func init() {
	once.Do(func() {
		file, _ := exec.LookPath(os.Args[0])
		path, _ := filepath.Abs(file)
		index := strings.LastIndex(path, string(os.PathSeparator))
		path = path[:index]
		index = strings.LastIndex(path, string(os.PathSeparator))
		path = path[:index]
		GetDeviceModels(path + "/")
	})
}

func GetDeviceModels(configPath string) (bool, map[string]DeviceModel) {

	//读取配置文件目录。
	yamlFiles, err := ioutil.ReadDir(configPath)
	if err != nil {
		Log.Error("Read config path ${root}/base/config/ err:{}", err.Error())
		return false, nil
	}
	//cycle read config yaml files transform iot-server type info
	for i := 0; i < len(yamlFiles); i++ {
		file := yamlFiles[i]
		if !file.IsDir() && strings.Contains(file.Name(), "dm-") {
			yamlFile, err := ioutil.ReadFile(configPath + file.Name())
			if err != nil {
				Log.Error("read yaml file err:{}", err.Error())
				return false, nil
			}
			deviceModel := getDeviceModel(file.Name(), yamlFile)
			DeviceModelMap[deviceModel.ID] = deviceModel
		}
	}
	return true, DeviceModelMap
}

//初始化设备类型
func getDeviceModel(filename string, bytes []byte) DeviceModel {
	deviceModel := DeviceModel{}
	err := yaml.Unmarshal(bytes, &deviceModel)
	if err != nil {
		panic("Read config file of iot-server-type err: " + filename + "" + err.Error())
	}
	//初始化SQL
	readDynamicSQL := DynamicSQL{}
	readDynamicSQL.TableName = "read_attr_" + deviceModel.ID
	readDynamicSQL.CreateTableSQL = ""
	//TODO 拼sql
	return deviceModel
}
