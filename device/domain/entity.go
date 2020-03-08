package domain

//
//import (
//	"errors"
//	"github.com/edwingeng/wuid/callback/wuid"
//	"gopkg.in/yaml.v2"
//	. "hui-iot/base/utils"
//	"io/ioutil"
//	"strings"
//	"time"
//)
//
//// config file path
//const configPath = "../config/devicetype/"
//
////The config file prefix of key
//type PrefixKey string
//
//const (
//	DataType     string = "data_type"
//	unit         string = "unit"
//	DefaultValue string = "default_value"
//)
//
//const (
//	PrefixkeyK PrefixKey = "k"
//	PrefixkeyI PrefixKey = "i"
//	PrefixkeyR PrefixKey = "r"
//	PrefixkeyC PrefixKey = "c"
//	PrefixkeyF PrefixKey = "f"
//)
//
//// save device type info
//var DeviceTypeMap = make(map[string]DeviceType)
//
//// base entity
//type BaseEntity struct {
//	ID         int64 //系统唯一标识
//	CreateTime time.Time
//	UpdateTime time.Time
//}
//
//var idUtil *wuid.WUID
//
//const id_tag = "iot"
//
////init base entity info
//func initIdUtil() {
//	wuid.WithSection(1)
//	idUtil = wuid.NewWUID(id_tag, nil)
//}
//
////init base info
//func (be *BaseEntity) InitBaseInfo() {
//	be.ID = idUtil.Next()
//	be.CreateTime = time.Now()
//	be.UpdateTime = time.Now()
//}
//
////动态生成表和对应的sql
//type DynamicSQL struct {
//	TableName      string //表名 ，格式：iot_[DeviceTypeID]_[PAttr|RAttr|CAttr|Func]
//	CreateTableSQL string //
//	DropTableSQL   string
//	InsertSQL      string
//	QuerySQL       string
//	UpdateSQL      string
//	DeleteSQL      string
//}
//
//// DeviceType
//type DeviceType struct {
//	ID       string       //唯一标识
//	InvarantDefine []AttrDefine //固定属性
//	ReadDefine  []AttrDefine //只读属性
//	ConfigDefine []AttrDefine //配置属性
//	Function []string     //功能函数
//	InvarantAttrSQL DynamicSQL   //固定属性动态sql
//	ReadAttrSQL DynamicSQL   //动态sql
//	ConfigAttrSQL DynamicSQL
//	FuncSQL  DynamicSQL
//}
//
////基本结构
//type KeyValue struct {
//	Key   string      //字段名称
//	Value interface{} //值
//}
//
////属性定义
//type AttrDefine struct {
//	KeyValue
//	DataType string // 字段类型：默认字符串
//	unit     string // 单位：默认值无
//}
//
//type ResultDTO struct {
//	Code int         `json:"code"` //Succ:1  or err code :<1
//	Data interface{} `json:"data"` //Result data
//	Msg  string      `json:"msg"`  //Succ or Err msg
//}
//
//func BuildSucc(result *ResultDTO) *ResultDTO {
//	result.Code = 1
//	if result.Msg == "" {
//		result.Msg = "succ"
//	}
//	return result
//}
//
//func BuildError(result *ResultDTO) *ResultDTO {
//	result.Code = 0
//	if result.Msg == "" {
//		result.Msg = "error"
//	}
//	return result
//}
//
////初始化配置文件
//func init() {
//	initIdUtil()      // Init Id generator.
//	initDeviceTypes() //Init device type config info.
//}
//
//func initDeviceTypes() {
//	//读取配置文件目录。
//	yamlFiles, err := ioutil.ReadDir(configPath)
//	if err != nil {
//		Log.Error("Read config path ${root}/base/config/ err:{}", err.Error())
//	}
//	//cycle read config yaml files transform device type info
//	for i := 0; i < len(yamlFiles); i++ {
//		file := yamlFiles[i]
//		if !file.IsDir() {
//			yamlFile, err := ioutil.ReadFile(configPath + file.Name())
//			if err != nil {
//				Log.Error("read yaml file err:{}", err.Error())
//			}
//			deviceType, err := initDeviceType(file.Name(), yamlFile)
//			if err != nil {
//				Log.Error("read yaml file err:{}", err.Error())
//			}
//			DeviceTypeMap[deviceType.ID] = deviceType
//		}
//	}
//}
//
////初始化设备类型
//func initDeviceType(filename string, bytes []byte) (DeviceType, error) {
//	Log.Debug("init yaml file name is ", filename)
//	errMsg := ""
//	deviceType := DeviceType{}
//	deviceType.ID = strings.Split(filename, ".")[0]
//	i := []AttrDefine{} //fixed attr
//	r := []AttrDefine{} //only read attr
//	c := []AttrDefine{} //config attr
//	f := []string{}     // function
//	configMap := make(map[string]interface{})
//	yaml.Unmarshal(bytes, configMap)
//	for k, v := range configMap {
//		prefix := strings.Split(k, "-")[0]
//		suffix := strings.Split(k, "-")[1]
//		switch prefix {
//		case string(PrefixkeyK):
//			//Not deal with
//		case string(PrefixkeyI):
//			i = append(i, NewAttrDefine(k, v))
//		case string(PrefixkeyR):
//			r = append(r, NewAttrDefine(k, v))
//		case string(PrefixkeyC):
//			c = append(c, NewAttrDefine(k, v))
//		case string(PrefixkeyF):
//			f = append(f, suffix)
//		default:
//			errMsg = "The yaml file format err,please check file [" + filename + "]"
//			//Stop server
//			Log.Panic(errMsg)
//		}
//	}
//	//Err of nil
//	if len(i)+len(r)+len(c)+len(f) == 0 {
//		errMsg = "The yaml file format err,please check file [" + filename + "]"
//		Log.Panic(errMsg)
//	}
//	deviceType.InvarantDefine = i
//	deviceType.ReadDefine = r
//	deviceType.ConfigDefine = c
//	deviceType.Function = f
//	//Is true
//	if len(errMsg) == 0 {
//		return deviceType, nil
//	}
//	//Is err
//	return deviceType, errors.New(errMsg)
//}
//
////创建
//func NewAttrDefine(key string, val interface{}) AttrDefine {
//	ad := AttrDefine{}
//	ad.Key = key
//	valMap, ok := val.(map[interface{}]interface{})
//	if ok {
//		//Get default value
//		value, ok := valMap[DefaultValue]
//		if ok {
//			ad.Value = value
//		}
//		//Get data type
//		dataType, ok := valMap[DataType]
//		if ok {
//			ad.DataType = dataType.(string)
//		}
//		//Get unit
//		unit, ok := valMap[unit]
//		if ok {
//			ad.unit = unit.(string)
//		}
//	}
//	return ad
//}
