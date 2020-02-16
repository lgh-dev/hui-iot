package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _IotDeviceTypeEventMgr struct {
	*_BaseMgr
}

// IotDeviceTypeEventMgr open func
func IotDeviceTypeEventMgr(db *gorm.DB) *_IotDeviceTypeEventMgr {
	if db == nil {
		panic(fmt.Errorf("IotDeviceTypeEventMgr need init by db"))
	}
	return &_IotDeviceTypeEventMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IotDeviceTypeEventMgr) GetTableName() string {
	return "iot_device_type_event"
}

// Get 获取
func (obj *_IotDeviceTypeEventMgr) Get() (result IotDeviceTypeEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IotDeviceTypeEventMgr) Gets() (results []*IotDeviceTypeEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID 主键
func (obj *_IotDeviceTypeEventMgr) WithID(ID int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithDeviceTypeID device_type_id获取 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceTypeEventMgr) WithDeviceTypeID(DeviceTypeID int64) Option {
	return optionFunc(func(o *options) { o.query["device_type_id"] = DeviceTypeID })
}

// WithEventName event_name获取 事件名称 事件名称
func (obj *_IotDeviceTypeEventMgr) WithEventName(EventName string) Option {
	return optionFunc(func(o *options) { o.query["event_name"] = EventName })
}

// WithEventLevel event_level获取 事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)
func (obj *_IotDeviceTypeEventMgr) WithEventLevel(EventLevel int) Option {
	return optionFunc(func(o *options) { o.query["event_level"] = EventLevel })
}

// GetByOption 功能选项模式获取
func (obj *_IotDeviceTypeEventMgr) GetByOption(opts ...Option) (result IotDeviceTypeEvent, err error) {
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
func (obj *_IotDeviceTypeEventMgr) GetByOptions(opts ...Option) (results []*IotDeviceTypeEvent, err error) {
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

// GetFromID 通过id获取内容 ID 主键
func (obj *_IotDeviceTypeEventMgr) GetFromID(ID int64) (result IotDeviceTypeEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 ID 主键
func (obj *_IotDeviceTypeEventMgr) GetBatchFromID(IDs []int64) (results []*IotDeviceTypeEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromDeviceTypeID 通过device_type_id获取内容 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceTypeEventMgr) GetFromDeviceTypeID(DeviceTypeID int64) (results []*IotDeviceTypeEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id = ?", DeviceTypeID).Find(&results).Error

	return
}

// GetBatchFromDeviceTypeID 批量唯一主键查找 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceTypeEventMgr) GetBatchFromDeviceTypeID(DeviceTypeIDs []int64) (results []*IotDeviceTypeEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id IN (?)", DeviceTypeIDs).Find(&results).Error

	return
}

// GetFromEventName 通过event_name获取内容 事件名称 事件名称
func (obj *_IotDeviceTypeEventMgr) GetFromEventName(EventName string) (results []*IotDeviceTypeEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_name = ?", EventName).Find(&results).Error

	return
}

// GetBatchFromEventName 批量唯一主键查找 事件名称 事件名称
func (obj *_IotDeviceTypeEventMgr) GetBatchFromEventName(EventNames []string) (results []*IotDeviceTypeEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_name IN (?)", EventNames).Find(&results).Error

	return
}

// GetFromEventLevel 通过event_level获取内容 事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)
func (obj *_IotDeviceTypeEventMgr) GetFromEventLevel(EventLevel int) (results []*IotDeviceTypeEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_level = ?", EventLevel).Find(&results).Error

	return
}

// GetBatchFromEventLevel 批量唯一主键查找 事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)
func (obj *_IotDeviceTypeEventMgr) GetBatchFromEventLevel(EventLevels []int) (results []*IotDeviceTypeEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_level IN (?)", EventLevels).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_IotDeviceTypeEventMgr) FetchByPrimaryKey(ID int64) (result IotDeviceTypeEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}