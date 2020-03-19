package iot_worker

import (
	"encoding/base64"
	"encoding/json"

	//_ "github.com/taosdata/TDengine/src/connector/go/src/taosSql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	sqlUrl = "/rest/sql"
)

//TDengine
type TDengineRestful struct {
	Ip       string `yaml:"ip"`       //IP
	Port     string `yaml:"port"`     //端口
	UserName string `yaml:"username"` //用户名
	Password string `yaml:"password"` //密码
	DBname   string `yaml:"dbname"`   //数据库名称
}

func NewTDengineRestful() *DataDB {
	once.Do(func() {
		//读取配置文件
		yamlFile, err := ioutil.ReadFile(configPath)
		if err != nil {
			panic("read config file err,path " + configPath)
		}
		td := TDengineRestful{}
		yaml.Unmarshal(yamlFile, &td)
		if db == nil {
			db = &td
		}
	})
	return &db
}

func getHTTPURI(td *TDengineRestful, uri string) string {
	return "http://" + td.Ip + ":" + td.Port + uri
}

func (td *TDengineRestful) CheckDBExists() bool {
	checkDBSql := "CREATE DATABASE IF NOT EXISTS " + td.DBname
	client := &http.Client{}
	req, err := http.NewRequest("POST", getHTTPURI(td, sqlUrl), strings.NewReader(checkDBSql))
	if err != nil {
		return false
	}
	req.Header.Set("Authorization", "Basic "+base64.URLEncoding.EncodeToString([]byte(td.UserName+":"+td.Password)))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return false
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	resultMap := make(map[string]string)
	json.Unmarshal(body, &resultMap)
	if resultMap["status"] == "succ" {
		return true
	} else {
		return false
	}
	return true
}

func (td *TDengineRestful) Insert(sql string) bool {
	return false
}

func (td *TDengineRestful) Get(sql string) interface{} {
	return nil
}

func (td *TDengineRestful) Query(sql string) []interface{} {
	return nil
}
