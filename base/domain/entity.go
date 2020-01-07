package domain

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	. "light-iot/base/utils"
	"strings"
)

// config file path
const configPath = "base/configfile/"

//The config file prefix of key
type PrefixKey string

const (
	DataType     string = "data_type"
	unit         string = "unit"
	DefaultValue string = "default_value"
)

const (
	PrefixkeyK PrefixKey = "k"
	PrefixkeyP PrefixKey = "p"
	PrefixkeyR PrefixKey = "r"
	PrefixkeyC PrefixKey = "c"
	PrefixkeyF PrefixKey = "f"
)

// save device type info
var DeviceTypeMap = make(map[string]DeviceType)

// device
type Device struct {
	ID           string                 //系统唯一标识
	UID          string                 `json:uid`  //可识别唯一标识
	Name         string                 `json:name` //可识别唯一标识
	P            map[string]interface{} //固定属性
	R            map[string]interface{} //只读属性
	C            map[string]interface{} //配置属性
	F            map[string]interface{} //功能函数
	DeviceTypeID string                 `json:deviceTypeID` //设备类型ID
}

// DeviceType
type DeviceType struct {
	ID string     //唯一标识
	P  []KeyValue //固定属性
	R  []KeyValue //只读属性
	C  []KeyValue //配置属性
	F  []string   //功能函数
}

//基本结构
type KeyValue struct {
	Key      string
	Val      interface{}
	DataType string
	unit     string
}

type ResultDTO struct {
	Code int         `json:"code"` //Succ or err code
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
	//读取配置文件目录。
	yamlFiles, err := ioutil.ReadDir(configPath)
	if err != nil {
		Log.Error("Read config path ${root}/base/configfile/ err:{}", err.Error())
	}
	//cycle read config yaml files transform device type info
	for i := 0; i < len(yamlFiles); i++ {
		file := yamlFiles[i]
		if !file.IsDir() {
			yamlFile, err := ioutil.ReadFile(configPath + file.Name())
			if err != nil {
				Log.Error("read yaml file err:{}", err.Error())
			}
			deviceType, err := initDeviceType(file.Name(), yamlFile)
			if err != nil {
				Log.Error("read yaml file err:{}", err.Error())
			}
			DeviceTypeMap[deviceType.ID] = deviceType
		}
	}
}

//初始化设备类型
func initDeviceType(filename string, bytes []byte) (DeviceType, error) {
	Log.Debug("init yaml file name is ", filename)
	errMsg := ""
	deviceType := DeviceType{}
	deviceType.ID = strings.Split(filename, ".")[0]
	p := []KeyValue{} //fixed attr
	r := []KeyValue{} //only read attr
	c := []KeyValue{} //config attr
	f := []string{}   // function
	configMap := make(map[string]interface{})
	yaml.Unmarshal(bytes, configMap)
	for k, v := range configMap {
		prefix := strings.Split(k, "-")[0]
		suffix := strings.Split(k, "-")[1]
		switch prefix {
		case string(PrefixkeyK):
			//Not deal with
		case string(PrefixkeyP):
			p = append(p, NewKeyValue(k, v))
		case string(PrefixkeyR):
			r = append(r, NewKeyValue(k, v))
		case string(PrefixkeyC):
			c = append(c, NewKeyValue(k, v))
		case string(PrefixkeyF):
			f = append(f, suffix)
		default:
			errMsg = "The yaml file format err,please check file [" + filename + "]"
			//Stop server
			Log.Panic(errMsg)
		}
	}
	//Err of nil
	if len(p)+len(r)+len(c)+len(f) == 0 {
		errMsg = "The yaml file format err,please check file [" + filename + "]"
		Log.Panic(errMsg)
	}
	deviceType.P = p
	deviceType.R = r
	deviceType.C = c
	deviceType.F = f
	//Is true
	if len(errMsg) == 0 {
		return deviceType, nil
	}
	//Is err
	return deviceType, errors.New(errMsg)
}

//创建基本结构
func NewKeyValue(key string, val interface{}) KeyValue {
	kv := KeyValue{}
	kv.Key = key
	valMap, ok := val.(map[interface{}]interface{})
	if ok {
		//Get default value
		value, ok := valMap[DefaultValue]
		if ok {
			kv.Val = value
		}
		//Get data type
		dataType, ok := valMap[DataType]
		if ok {
			kv.DataType = dataType.(string)
		}
		//Get unit
		unit, ok := valMap[unit]
		if ok {
			kv.unit = unit.(string)
		}
	}
	return kv
}
