package dao

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	. "hui-iot/base/domain"
	"hui-iot/device/api/dto"

	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/10 0:55
 * @Desc:
 */

type DBOperation interface {
	add(device *Device) error
	query(id dto.EntityID) *Device
	update(device *Device) *Device
	delete(id dto.EntityID) *Device
}

type DBInfo struct {
	dbType     string   //Database type such as mongodb or mysql
	hosts      []string //Database connect hosts such as 127.0.0.1:27017
	database   string   //Database
	username   string   //Database username
	password   string   //Database password
	collection string   //Database collection
}

var (
	mongo = DBInfo{"mongodb", []string{"mongodb://127.0.0.1:27017"}, "iot_device", "lgh", "lgh_)(*&^)", "device"}
)

var GlobalMgoSession *mgo.Session

//mongodb连接常量
const (
	key_id     = "_id"
	key_status = "status"
	bson_set   = "$set"
	timeFormat = "2006-01-02 15:04:05"
)

//init db client
func init() {
	info := &mgo.DialInfo{
		Addrs:    mongo.hosts,
		Timeout:  time.Second * 5,
		Database: mongo.database,
		Username: mongo.username,
		Password: mongo.password,
	}
	globalMgoSession, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}
	GlobalMgoSession = globalMgoSession
	GlobalMgoSession.SetMode(mgo.Monotonic, true)
	globalMgoSession.SetPoolLimit(1000)
	log.Println("初始化连接数据库成功！")
}

func CloneSession() *mgo.Session {
	return GlobalMgoSession.Clone()
}

func (db *DBInfo) add(device *Device) error {
	session := CloneSession()
	defer session.Close()
	c := session.DB(mongo.database).C(mongo.collection)
	return c.Insert(device)
}
func (db *DBInfo) query(id dto.EntityID) *Device {
	session := CloneSession()
	defer session.Close()
	var device = Device{}
	c := session.DB(mongo.database).C(mongo.collection)
	c.Find(bson.D{{"_id", id}, {"status", 1}}).One(&device)
	return &device
}

//TODO
func (db *DBInfo) update(device *Device) error {
	session := CloneSession()
	defer session.Close()
	c := session.DB(mongo.database).C(mongo.collection)
	return c.Update(device.ID, bson.M{"$set": bson.M{"status": 0, "updateTime": time.Now().Format(timeFormat)}})
}

func (db *DBInfo) delete(id dto.EntityID) error {
	session := CloneSession()
	defer session.Close()
	c := session.DB(mongo.database).C(mongo.collection)
	return c.Update(id, bson.M{bson_set: bson.M{key_status: 0, "updateTime": time.Now()}})
}
