package config

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"strings"
)

// save iot-server types info
var DeviceModelMap = make(map[string]map[string]interface{})

func InitDeviceModels(configPath string) (bool, map[string]map[string]interface{}) {
	//
	//	//读取配置文件目录。
	deviceModelConfigFiles, err := ioutil.ReadDir(configPath)
	if err != nil {
		log.Printf("Read config path ${root}/conf/ err:%s", err.Error())
		return false, nil
	}
	//cycle read config yaml files transform iot-server type info
	for i := 0; i < len(deviceModelConfigFiles); i++ {
		file := deviceModelConfigFiles[i]
		if !file.IsDir() && strings.Contains(file.Name(), "dm-") {
			v := viper.New()
			v.SetConfigName(file.Name()) // name of config file (without extension)
			v.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
			v.AddConfigPath(configPath + "/")
			v.ReadInConfig()
			deviceModelName := strings.TrimPrefix(file.Name(), "dm-")
			deviceModelName = strings.TrimSuffix(deviceModelName, ".yaml")
			deviceModel := v.GetStringMap(deviceModelName)
			DeviceModelMap[deviceModelName] = deviceModel
		}
	}
	return true, DeviceModelMap
}
