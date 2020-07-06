package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"hui-iot/iot-server/domain"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

var Conf *viper.Viper

var once sync.Once // 单例工具。

var AppYaml *domain.AppYaml

const projectName = "hui-iot"

func init() {
	InitAllConfig() //加载所有配置
}

func InitAllConfig() {
	confPath := GetPath() + "/conf"
	InitConfig(confPath)                         //初始化基本配置
	InitDeviceModels(confPath + "/device_model") //初始化设备模型配置
	WatchConfig()
}

func GetPath() string {
	path := os.Getenv("HUI_IOT_HOME")
	if path != "" {
		return path
	}
	log.Printf("HUI_IOT_HOME is %s\n", path)
	//获取当前文件的路径 /Users/lgh/Documents/leavemsg/config/config.go
	_, filename, _, _ := runtime.Caller(0)
	//获取项目目录 /Users/lgh/Documents/conf-demo
	filename = strings.Split(filename, projectName)[0] + projectName
	return filename
}

// 获取全局变量 ,单例模式。
func InitConfig(path string) *viper.Viper {
	once.Do(func() {
		if Conf == nil {
			Conf = ReadConfigFile(path, "iot-worker", "yaml")
		}
	})
	return Conf
}

//读取配置文件。
func ReadConfigFile(path string, fileName string, configType string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(fileName)   // name of config file (without extension)
	v.SetConfigType(configType) // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath(path)
	err := v.ReadInConfig()
	yamlfile, err := ioutil.ReadFile(path + "/" + fileName + "." + configType)
	yaml.Unmarshal(yamlfile, &AppYaml) // Find and read the config file
	if err != nil {                    // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return v
}
func WatchConfig() {
	Conf.WatchConfig()
	log.Println("start to watch config file!")
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s\n", e.Name)
	})
}
