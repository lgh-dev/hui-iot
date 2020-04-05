package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _IotDeviceEventMgr struct {
	*_BaseMgr
}

// IotDeviceEventMgr open func
func IotDeviceEventMgr(db *gorm.DB) *_IotDeviceEventMgr {
	if db == nil {
		panic(fmt.Errorf("IotDeviceEventMgr need init by db"))
	}
	return &_IotDeviceEventMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IotDeviceEventMgr) GetTableName() string {
	return "iot_device_event"
}

// Get 获取
func (obj *_IotDeviceEventMgr) Get() (result IotDeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IotDeviceEventMgr) Gets() (results []*IotDeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithTime time获取 落库时间 时序数据库的主键
func (obj *_IotDeviceEventMgr) WithTime(Time int64) Option {
	return optionFunc(func(o *options) { o.query["time"] = Time })
}

// WithEventTime event_time获取 事件发送时间
func (obj *_IotDeviceEventMgr) WithEventTime(EventTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["event_time"] = EventTime })
}

// WithDeviceModelEventID device_type_event_id获取 设备类型事件定义ID-外键 设备类型事件定义ID-外键
func (obj *_IotDeviceEventMgr) WithDeviceModelEventID(DeviceModelEventID int64) Option {
	return optionFunc(func(o *options) { o.query["device_type_event_id"] = DeviceModelEventID })
}

// WithEventInfoMap event_info_map获取 事件信息 事件信息，json，{key:value}结构
func (obj *_IotDeviceEventMgr) WithEventInfoMap(EventInfoMap string) Option {
	return optionFunc(func(o *options) { o.query["event_info_map"] = EventInfoMap })
}

// WithEventLevel event_level获取 事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)
func (obj *_IotDeviceEventMgr) WithEventLevel(EventLevel int) Option {
	return optionFunc(func(o *options) { o.query["event_level"] = EventLevel })
}

// WithEventStatus event_status获取 事件状态 事件状态（告警1，结束0）
func (obj *_IotDeviceEventMgr) WithEventStatus(EventStatus string) Option {
	return optionFunc(func(o *options) { o.query["event_status"] = EventStatus })
}

// GetByOption 功能选项模式获取
func (obj *_IotDeviceEventMgr) GetByOption(opts ...Option) (result IotDeviceEvent, err error) {
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
func (obj *_IotDeviceEventMgr) GetByOptions(opts ...Option) (results []*IotDeviceEvent, err error) {
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
func (obj *_IotDeviceEventMgr) GetFromTime(Time int64) (result IotDeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time = ?", Time).Find(&result).Error

	return
}

// GetBatchFromTime 批量唯一主键查找 落库时间 时序数据库的主键
func (obj *_IotDeviceEventMgr) GetBatchFromTime(Times []int64) (results []*IotDeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time IN (?)", Times).Find(&results).Error

	return
}

// GetFromEventTime 通过event_time获取内容 事件发送时间
func (obj *_IotDeviceEventMgr) GetFromEventTime(EventTime time.Time) (results []*IotDeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_time = ?", EventTime).Find(&results).Error

	return
}

// GetBatchFromEventTime 批量唯一主键查找 事件发送时间
func (obj *_IotDeviceEventMgr) GetBatchFromEventTime(EventTimes []time.Time) (results []*IotDeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_time IN (?)", EventTimes).Find(&results).Error

	return
}

// GetFromDeviceModelEventID 通过device_type_event_id获取内容 设备类型事件定义ID-外键 设备类型事件定义ID-外键
func (obj *_IotDeviceEventMgr) GetFromDeviceModelEventID(DeviceModelEventID int64) (results []*IotDeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_event_id = ?", DeviceModelEventID).Find(&results).Error

	return
}

// GetBatchFromDeviceModelEventID 批量唯一主键查找 设备类型事件定义ID-外键 设备类型事件定义ID-外键
func (obj *_IotDeviceEventMgr) GetBatchFromDeviceModelEventID(DeviceModelEventIDs []int64) (results []*IotDeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_event_id IN (?)", DeviceModelEventIDs).Find(&results).Error

	return
}

// GetFromEventInfoMap 通过event_info_map获取内容 事件信息 事件信息，json，{key:value}结构
func (obj *_IotDeviceEventMgr) GetFromEventInfoMap(EventInfoMap string) (results []*IotDeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_info_map = ?", EventInfoMap).Find(&results).Error

	return
}

// GetBatchFromEventInfoMap 批量唯一主键查找 事件信息 事件信息，json，{key:value}结构
func (obj *_IotDeviceEventMgr) GetBatchFromEventInfoMap(EventInfoMaps []string) (results []*IotDeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_info_map IN (?)", EventInfoMaps).Find(&results).Error

	return
}

// GetFromEventLevel 通过event_level获取内容 事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)
func (obj *_IotDeviceEventMgr) GetFromEventLevel(EventLevel int) (results []*IotDeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_level = ?", EventLevel).Find(&results).Error

	return
}

// GetBatchFromEventLevel 批量唯一主键查找 事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)
func (obj *_IotDeviceEventMgr) GetBatchFromEventLevel(EventLevels []int) (results []*IotDeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_level IN (?)", EventLevels).Find(&results).Error

	return
}

// GetFromEventStatus 通过event_status获取内容 事件状态 事件状态（告警1，结束0）
func (obj *_IotDeviceEventMgr) GetFromEventStatus(EventStatus string) (results []*IotDeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_status = ?", EventStatus).Find(&results).Error

	return
}

// GetBatchFromEventStatus 批量唯一主键查找 事件状态 事件状态（告警1，结束0）
func (obj *_IotDeviceEventMgr) GetBatchFromEventStatus(EventStatuss []string) (results []*IotDeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_status IN (?)", EventStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_IotDeviceEventMgr) FetchByPrimaryKey(Time int64) (result IotDeviceEvent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time = ?", Time).Find(&result).Error

	return
}
