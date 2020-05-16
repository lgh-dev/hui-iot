package db

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"hui-iot/iot-worker/config"

	"io/ioutil"
	"net/http"
	"strings"
)

var (
	sqlUrl = "/rest/sql"
	td     *TDengineRestful
	client = &http.Client{}
)

//TDengine
type TDengineRestful struct {
	Ip       string `yaml:"ip"`       //IP
	Port     string `yaml:"port"`     //端口
	UserName string `yaml:"username"` //用户名
	Password string `yaml:"password"` //密码
	DBname   string `yaml:"dbname"`   //数据库名称
}

func NewTDengineRestful() *TDengineRestful {
	once.Do(func() {
		if td == nil {
			td = &TDengineRestful{Ip: config.Conf.GetString("db.TDengine.ip"),
				Port:     config.Conf.GetString("db.TDengine.port"),
				UserName: config.Conf.GetString("db.TDengine.username"),
				Password: config.Conf.GetString("db.TDengine.password"),
				DBname:   config.Conf.GetString("db.TDengine.dbname")}
		}
	})
	return td
}

func getHTTPURI(td *TDengineRestful, uri string) string {
	return "http://" + td.Ip + ":" + td.Port + uri
}

func (td *TDengineRestful) CheckDBExists() bool {
	checkDBSql := "CREATE DATABASE IF NOT EXISTS " + td.DBname
	return execSQL(td, checkDBSql) == nil
}

func execSQL(td *TDengineRestful, sql string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", getHTTPURI(td, sqlUrl), strings.NewReader(sql))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Basic "+base64.URLEncoding.EncodeToString([]byte(td.UserName+":"+td.Password)))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resultMap := make(map[string]string)
	json.Unmarshal(body, &resultMap)
	if resultMap["status"] == "succ" {
		return nil
	} else {
		return errors.New(string(body))
	}
	return err
}

func (td *TDengineRestful) Insert(sql string) error {
	return execSQL(td, sql)
}

func (td *TDengineRestful) Get(sql string) interface{} {
	return nil
}

func (td *TDengineRestful) Query(sql string) []interface{} {
	return nil
}
