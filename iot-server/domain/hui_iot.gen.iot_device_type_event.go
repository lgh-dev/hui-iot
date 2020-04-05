package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _IotDeviceModelEventMgr struct {
	*_BaseMgr
}

// IotDeviceModelEventMgr open func
func IotDeviceModelEventMgr(db *gorm.DB) *_IotDeviceModelEventMgr {
	if db == nil {
		panic(fmt.Errorf("IotDeviceModelEventMgr need init by db"))
	}
	return &_IotDeviceModelEventMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IotDeviceModelEventMgr) GetTableName() string {
	return "iot_device_type_event"
}

// Get 获取
func (obj *_IotDeviceModelEventMgr) Get() (result IotDeviceModelEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IotDeviceModelEventMgr) Gets() (results []*IotDeviceModelEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID 主键
func (obj *_IotDeviceModelEventMgr) WithID(ID int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithDeviceModelID device_type_id获取 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceModelEventMgr) WithDeviceModelID(DeviceModelID int64) Option {
	return optionFunc(func(o *options) { o.query["device_type_id"] = DeviceModelID })
}

// WithEventName event_name获取 事件名称 事件名称
func (obj *_IotDeviceModelEventMgr) WithEventName(EventName string) Option {
	return optionFunc(func(o *options) { o.query["event_name"] = EventName })
}

// WithEventLevel event_level获取 事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)
func (obj *_IotDeviceModelEventMgr) WithEventLevel(EventLevel int) Option {
	return optionFunc(func(o *options) { o.query["event_level"] = EventLevel })
}

// GetByOption 功能选项模式获取
func (obj *_IotDeviceModelEventMgr) GetByOption(opts ...Option) (result IotDeviceModelEvent, err error) {
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
func (obj *_IotDeviceModelEventMgr) GetByOptions(opts ...Option) (results []*IotDeviceModelEvent, err error) {
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
func (obj *_IotDeviceModelEventMgr) GetFromID(ID int64) (result IotDeviceModelEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 ID 主键
func (obj *_IotDeviceModelEventMgr) GetBatchFromID(IDs []int64) (results []*IotDeviceModelEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromDeviceModelID 通过device_type_id获取内容 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceModelEventMgr) GetFromDeviceModelID(DeviceModelID int64) (results []*IotDeviceModelEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id = ?", DeviceModelID).Find(&results).Error

	return
}

// GetBatchFromDeviceModelID 批量唯一主键查找 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceModelEventMgr) GetBatchFromDeviceModelID(DeviceModelIDs []int64) (results []*IotDeviceModelEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id IN (?)", DeviceModelIDs).Find(&results).Error

	return
}

// GetFromEventName 通过event_name获取内容 事件名称 事件名称
func (obj *_IotDeviceModelEventMgr) GetFromEventName(EventName string) (results []*IotDeviceModelEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_name = ?", EventName).Find(&results).Error

	return
}

// GetBatchFromEventName 批量唯一主键查找 事件名称 事件名称
func (obj *_IotDeviceModelEventMgr) GetBatchFromEventName(EventNames []string) (results []*IotDeviceModelEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_name IN (?)", EventNames).Find(&results).Error

	return
}

// GetFromEventLevel 通过event_level获取内容 事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)
func (obj *_IotDeviceModelEventMgr) GetFromEventLevel(EventLevel int) (results []*IotDeviceModelEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_level = ?", EventLevel).Find(&results).Error

	return
}

// GetBatchFromEventLevel 批量唯一主键查找 事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)
func (obj *_IotDeviceModelEventMgr) GetBatchFromEventLevel(EventLevels []int) (results []*IotDeviceModelEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_level IN (?)", EventLevels).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_IotDeviceModelEventMgr) FetchByPrimaryKey(ID int64) (result IotDeviceModelEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}
