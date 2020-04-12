package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"runtime"
	"strings"
	"sync"
)

var Conf *viper.Viper

var once sync.Once // 单例工具。

const projectName = "hui-iot"

func init() {
	InitConfig(GetPath() + "/conf")
	WatchConfig()
}

// 获取全局变量 ,单例模式。
func InitConfig(path string) *viper.Viper {
	once.Do(func() {
		if Conf == nil {
			Conf = ReadConfigFile(path, "app", "yaml")
		}
	})
	return Conf
}

func GetPath() string {
	//获取当前文件的路径 /Users/lgh/Documents/leavemsg/config/config.go
	_, filename, _, _ := runtime.Caller(0)
	//获取项目目录 /Users/lgh/Documents/conf-demo
	filename = strings.Split(filename, projectName)[0] + projectName
	return filename
}

//读取配置文件。
func ReadConfigFile(path string, fileName string, configType string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(fileName)   // name of config file (without extension)
	v.SetConfigType(configType) // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath(path)
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Print(v.AllKeys())
	return v
}
func WatchConfig() {
	Conf.WatchConfig()
	log.Println("start to watch config file!")
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
	})
}
