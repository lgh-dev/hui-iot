package modules

import (
	"github.com/globalsign/mgo"
	"hui-iot/iot-server/config"
	"log"
	"sync"
	"time"
)

//mongodb连接常量
const timeFormat = "2006-01-02 15:04:05"

//添加参数对象
type AddBody struct {
	Uid           string    `json:"uid"`
	Msg           string    `json:"msg"`
	ClientIp      string    `json:"clientIp"`
	Status        int8      `json:"status"` //状态，1是正常。0是删除。
	CreateTime    string    `json:"createTime"`
	UpdateTime    string    `json:"updateTime"`
	EffectiveTime time.Time `json:"effectiveTime"`
}

var GSession *mgo.Session

var dbInfo *mgo.DialInfo

//开发环境配置
const (
	collection = "msg"
)

var once sync.Once

func initOnce() {
	once.Do(func() {
		dbInfo = &mgo.DialInfo{
			Addrs:    []string{config.Conf.GetString("db.mongo.ip") + ":" + config.Conf.GetString("db.mongo.port")},
			Timeout:  time.Second * 5,
			Database: config.Conf.GetString("db.mongo.dbname"),
			Username: config.Conf.GetString("db.mongo.username"),
			Password: config.Conf.GetString("db.mongo.password"),
		}
		gSession, err := mgo.DialWithInfo(dbInfo)
		if err != nil {
			log.Printf("连接不到mongodb")
			panic(err)
		}
		GSession = gSession
		GSession.SetMode(mgo.Monotonic, true)
		GSession.SetPoolLimit(1000)
		log.Println("初始化连接数据库成功！")
	})

}

func CloneSession() *mgo.Session {
	if GSession == nil {
		initOnce()
	}
	return GSession.Clone()
}

//插入数据
func Insert(body *AddBody) {
	session := CloneSession()
	defer session.Close()
	c := session.DB(dbInfo.Database).C(collection)
	body.Status = 1
	body.CreateTime = time.Now().Format(timeFormat)
	body.UpdateTime = body.CreateTime
	c.Insert(body)
}
