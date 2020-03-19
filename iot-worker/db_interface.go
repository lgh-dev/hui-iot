package iot_worker

import "sync"

//数据服务的DB对象
type DataDB interface {
	CheckDBExists() bool            // 查询数据库是否存在
	Insert(sql string) bool         //插入
	Get(sql string) interface{}     //查询单个对象
	Query(sql string) []interface{} //查询多个对象
}

var (
	once sync.Once
	db   DataDB
)

const configPath = "config/db.yaml"
