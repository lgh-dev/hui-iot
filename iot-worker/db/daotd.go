package db

import (
	"database/sql"
	"fmt"
	_ "github.com/taosdata/driver-go/taosSql"
	"hui-iot/iot-worker/config"
	"hui-iot/iot-worker/utils"
)

//TDengine
type TDengine struct {
	DriverName string //驱动名称
	UserName   string //数据库用户
	Password   string //数据库密码
	Ip         string //IP
	Port       string //端口
	DBname     string //数据库名称
	Connect    *sql.DB
}

var tde *TDengine

const taosDriverName = "taosSql"

func NewTDengine() *TDengine {
	once.Do(func() {
		if tde == nil {
			tde = &TDengine{
				DriverName: taosDriverName,
				Ip:         config.Conf.GetString("db.TDengine.ip"),
				Port:       config.Conf.GetString("db.TDengine.port"),
				UserName:   config.Conf.GetString("db.TDengine.username"),
				Password:   config.Conf.GetString("db.TDengine.password"),
				DBname:     config.Conf.GetString("db.TDengine.dbname")}
			tde.Connect = tde.getDB()
		}

	})
	return tde
}

func (td *TDengine) CheckDBExists() bool {
	db := td.getDB()
	defer db.Close()
	//准备创建数据库语句
	sqlcmd := fmt.Sprintf("create database if not exists %s", td.DBname)
	//执行语句并检查错误
	_, err := db.Exec(sqlcmd)
	if err != nil {
		fmt.Printf("Create database error: %s\n", err)
		return false
	}
	return true
}

func (td *TDengine) getDB() *sql.DB {
	db, err := sql.Open(td.DriverName, td.UserName+":"+td.Password+"@/tcp("+td.Ip+":0"+")/"+td.DBname)
	if err != nil {
		fmt.Printf("Open database error: %s\n", err)
		return nil
	}
	return db
}

func (td *TDengine) Insert(sql *string) error {
	//db := td.getDB()
	//defer db.Close()
	//st := time.Now().Nanosecond()
	// insert data
	_, err := td.Connect.Exec(*sql)
	utils.CheckErrorLn("save to tdengine err:%s", err)
	return err
}

func (td *TDengine) Get(sql string) interface{} {
	return nil
}

func (td *TDengine) Query(sql string) []interface{} {
	return nil
}
