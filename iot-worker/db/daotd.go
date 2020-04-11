package db

import (
	"database/sql"
	"fmt"
	//_ "github.com/taosdata/TDengine/src/connector/go/src/taosSql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//TDengine
type TDengine struct {
	DriverName string //驱动名称
	DBuser     string //数据库用户
	DBpassword string //数据库密码
	Ip         string //IP
	Port       string //端口
	DBname     string //数据库名称
}

func NewTDengine() *DataDB {
	once.Do(func() {
		//读取配置文件
		yamlFile, err := ioutil.ReadFile(configPath)
		if err != nil {
			panic("read config file err,path " + configPath)
		}
		td := TDengine{}
		yaml.Unmarshal(yamlFile, &td)
		if db == nil {
			db = &td
		}
	})
	return &db
}

func (td *TDengine) CheckDBExists() bool {
	db, err := sql.Open(td.DriverName, td.DBuser+":"+td.DBpassword+"@/tcp("+td.Ip+":0"+")/")
	if err != nil {
		fmt.Printf("Open database error: %s\n", err)
		return false
	}
	defer db.Close()
	//准备创建数据库语句
	sqlcmd := fmt.Sprintf("create database if not exists %s", td.DBname)
	//执行语句并检查错误
	_, err = db.Exec(sqlcmd)
	if err != nil {
		fmt.Printf("Create database error: %s\n", err)
		return false
	}
	return true
}

func (td *TDengine) Insert(sql string) bool {
	return false
}

func (td *TDengine) Get(sql string) interface{} {
	return nil
}

func (td *TDengine) Query(sql string) []interface{} {
	return nil
}
