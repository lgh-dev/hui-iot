package config

import (
	"gopkg.in/yaml.v2"
	"hui-iot/iot-server/common"
	"hui-iot/iot-server/domain"
	"io/ioutil"
	"log"
)

// save iot-server types info
var AppDefineMap = make(map[string]domain.ApplyDefine)

func InitAppConfigs(configPath string) (bool, map[string]domain.ApplyDefine) {
	if len(AppYaml.Apples) > 0 {
		for _, apply := range AppYaml.Apples {
			if apply.Enable == true {
				appConfigName := configPath + "/apply-" + apply.AppKey + ".yaml"
				//	//读取配置文件目录。
				yamlFiles, err := ioutil.ReadFile(appConfigName)
				if err != nil {
					common.Log.Error("Read config path ${root}/base/config/ err:{}", err.Error())
					return false, nil
				}
				appDefine := readAppConfig(appConfigName, yamlFiles)
				if apply.Version != appDefine.Version {
					log.Panicf("Read %s err: version %s not exists ", appConfigName, apply.Version)
				}
				AppDefineMap[appDefine.AppKey] = appDefine
			}

		}
	}
	return true, AppDefineMap
}

//初始化设备类型
func readAppConfig(filename string, bytes []byte) domain.ApplyDefine {
	applyDefine := domain.ApplyDefine{}
	err := yaml.Unmarshal(bytes, &applyDefine)
	if err != nil {
		panic("Read config file of apply-xxx.yaml err: " + filename + "" + err.Error())
	}
	return applyDefine
}
