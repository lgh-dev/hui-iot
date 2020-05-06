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

type ICurdService interface {
	Add(Entity interface{}) (string, error)
	FindByID(id string) interface{}
	FindByIDs(ids []string) interface{}
	FindByPage(condition interface{}, page Page) interface{}
	Update(Entity interface{}) error
	Delete(ids []string) error
}
type Page struct {
	PageNumber int //
	PageSize   int
	Sort       map[string]int //排序，key是字段名，value=1是顺序，value=-1是倒序
}

type CurdService struct {
}

func getDBName() string {
	return config.Conf.GetString("db.mongo.dbname")
}

func getCollection(session *mgo.Session) *mgo.Collection {
	c := session.DB(getDBName()).C(collection)
	return c
}
