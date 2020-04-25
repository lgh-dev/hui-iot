package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _DeviceEventMgr struct {
	*_BaseMgr
}

// DeviceEventMgr open func
func DeviceEventMgr(db *gorm.DB) *_DeviceEventMgr {
	if db == nil {
		panic(fmt.Errorf("DeviceEventMgr need init by db"))
	}
	return &_DeviceEventMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DeviceEventMgr) GetTableName() string {
	return "device_event"
}

// Get 获取
func (obj *_DeviceEventMgr) Get() (result DeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DeviceEventMgr) Gets() (results []*DeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithTime time获取 落库时间 时序数据库的主键
func (obj *_DeviceEventMgr) WithTime(Time int64) Option {
	return optionFunc(func(o *options) { o.query["time"] = Time })
}

// WithSendTime send_time获取 事件发送时间
func (obj *_DeviceEventMgr) WithSendTime(SendTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["send_time"] = SendTime })
}

// WithUkey ukey获取 事件定义标识
func (obj *_DeviceEventMgr) WithUkey(Ukey int64) Option {
	return optionFunc(func(o *options) { o.query["ukey"] = Ukey })
}

// WithInfo info获取 事件信息
func (obj *_DeviceEventMgr) WithInfo(Info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = Info })
}

// WithLevel level获取 事件级别(信息INFO、告警ALARM、故障ERROR)
func (obj *_DeviceEventMgr) WithLevel(Level string) Option {
	return optionFunc(func(o *options) { o.query["level"] = Level })
}

// WithStatus status获取 事件状态 事件状态（告警ALARM，结束OVER）
func (obj *_DeviceEventMgr) WithStatus(Status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = Status })
}

// GetByOption 功能选项模式获取
func (obj *_DeviceEventMgr) GetByOption(opts ...Option) (result DeviceEvent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DeviceEventMgr) GetByOptions(opts ...Option) (results []*DeviceEvent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromTime 通过time获取内容 落库时间 时序数据库的主键
func (obj *_DeviceEventMgr) GetFromTime(Time int64) (result DeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time = ?", Time).Find(&result).Error

	return
}

// GetBatchFromTime 批量唯一主键查找 落库时间 时序数据库的主键
func (obj *_DeviceEventMgr) GetBatchFromTime(Times []int64) (results []*DeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time IN (?)", Times).Find(&results).Error

	return
}

// GetFromSendTime 通过send_time获取内容 事件发送时间
func (obj *_DeviceEventMgr) GetFromSendTime(SendTime time.Time) (results []*DeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("send_time = ?", SendTime).Find(&results).Error

	return
}

// GetBatchFromSendTime 批量唯一主键查找 事件发送时间
func (obj *_DeviceEventMgr) GetBatchFromSendTime(SendTimes []time.Time) (results []*DeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("send_time IN (?)", SendTimes).Find(&results).Error

	return
}

// GetFromUkey 通过ukey获取内容 事件定义标识
func (obj *_DeviceEventMgr) GetFromUkey(Ukey int64) (results []*DeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ukey = ?", Ukey).Find(&results).Error

	return
}

// GetBatchFromUkey 批量唯一主键查找 事件定义标识
func (obj *_DeviceEventMgr) GetBatchFromUkey(Ukeys []int64) (results []*DeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ukey IN (?)", Ukeys).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容 事件信息
func (obj *_DeviceEventMgr) GetFromInfo(Info string) (results []*DeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("info = ?", Info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量唯一主键查找 事件信息
func (obj *_DeviceEventMgr) GetBatchFromInfo(Infos []string) (results []*DeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("info IN (?)", Infos).Find(&results).Error

	return
}

// GetFromLevel 通过level获取内容 事件级别(信息INFO、告警ALARM、故障ERROR)
func (obj *_DeviceEventMgr) GetFromLevel(Level string) (results []*DeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("level = ?", Level).Find(&results).Error

	return
}

// GetBatchFromLevel 批量唯一主键查找 事件级别(信息INFO、告警ALARM、故障ERROR)
func (obj *_DeviceEventMgr) GetBatchFromLevel(Levels []string) (results []*DeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("level IN (?)", Levels).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 事件状态 事件状态（告警ALARM，结束OVER）
func (obj *_DeviceEventMgr) GetFromStatus(Status string) (results []*DeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", Status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 事件状态 事件状态（告警ALARM，结束OVER）
func (obj *_DeviceEventMgr) GetBatchFromStatus(Statuss []string) (results []*DeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", Statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_DeviceEventMgr) FetchByPrimaryKey(Time int64) (result DeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time = ?", Time).Find(&result).Error

	return
}
