package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
	"strings"
)

//初始化配置文件
func main() {

	yamlFiles, err := ioutil.ReadDir("base/configfile/")

	for i := 0; i < len(yamlFiles); i++ {
		file := yamlFiles[i]
		if !file.IsDir() {
			yamlFile, err := ioutil.ReadFile(file.Name())
			if err != nil {
				fmt.Printf("read yaml file err:{}", err.Error())
			}
			deviceType, err := initDeviceType(file.Name(), yamlFile)
			fmt.Printf("deviceType[{}]：{}", i, deviceType)
		}
	}
}

//初始化设备类型
func initDeviceType(filename string, bytes []byte) (DeviceType, error) {
	deviceType := DeviceType{}
	deviceType.id = filename
	p:=[]KeyValue{}
	r:=[]KeyValue{}
	c:=[]KeyValue{}
	f:=[]string{}
	configMap := new(map[string]interface{})
	yaml.Unmarshal(bytes, configMap)
	for k, v := range configMap {
		prefix:=strings.Split(k, "-")[0]
		suffix:=strings.Split(k, "-")[1]
		switch prefix {
		case "k":
		case "p":
			kv:=KeyValue{}
			kv.key=suffix
			kv.val=v

			kv.dataType==nil?configMap[k]
			p = append(p, KeyValue{key:suffix,val:v,dataType:configMap[k]["dataType"]})
		case "r":
		case "c":
		case "f":

		}
		strings.HasPrefix(k, "k-")
	}
	//fmt.Printf("yamlfile",testMap)
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v ", err)
	}

}

// device
type Device struct {
	id         string                 //唯一标识
	r          map[string]interface{} //只读属性
	c          map[string]interface{} //配置属性
	deviceType DeviceType             //设备类型
}

// device
type DeviceType struct {
	id string     //唯一标识
	p  []KeyValue //固定属性
	r  []KeyValue //只读属性
	c  []KeyValue //配置属性
	f  []string   //功能函数
}

//基本结构
type KeyValue struct {
	key      string
	val      interface{}
	dataType string
	unit     string
}
