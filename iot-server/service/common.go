package service

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/globalsign/mgo"
	"hui-iot/iot-server/config"
	"sync"
)

var once sync.Once

const collection = "device"

const (
	NONACTIVE = "NOTACTIVE" //未激活
	ONLINE    = "ONLINE"    //在线
	OFLINE    = "OFLINE"    //离线
)

func getDBName() string {
	return config.Conf.GetString("db.mongo.dbname")
}

func getCollection(session *mgo.Session) *mgo.Collection {
	c := session.DB(getDBName()).C(collection)
	return c
}
