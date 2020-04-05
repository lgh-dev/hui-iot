package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _IotDeviceEventStatusMgr struct {
	*_BaseMgr
}

// IotDeviceEventStatusMgr open func
func IotDeviceEventStatusMgr(db *gorm.DB) *_IotDeviceEventStatusMgr {
	if db == nil {
		panic(fmt.Errorf("IotDeviceEventStatusMgr need init by db"))
	}
	return &_IotDeviceEventStatusMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IotDeviceEventStatusMgr) GetTableName() string {
	return "iot_device_event_status"
}

// Get 获取
func (obj *_IotDeviceEventStatusMgr) Get() (result IotDeviceEventStatus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IotDeviceEventStatusMgr) Gets() (results []*IotDeviceEventStatus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithTime time获取 落库时间 时序数据库的主键
func (obj *_IotDeviceEventStatusMgr) WithTime(Time int64) Option {
	return optionFunc(func(o *options) { o.query["time"] = Time })
}

// WithEventTime event_time获取 事件发生时间 事件发生时间，如果上报为空，取服务器时间。
func (obj *_IotDeviceEventStatusMgr) WithEventTime(EventTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["event_time"] = EventTime })
}

// WithDeviceModelEventID device_type_event_id获取 设备类型事件定义ID-外键 设备类型事件定义ID-外键
func (obj *_IotDeviceEventStatusMgr) WithDeviceModelEventID(DeviceModelEventID int64) Option {
	return optionFunc(func(o *options) { o.query["device_type_event_id"] = DeviceModelEventID })
}

// WithEventInfoMap event_info_map获取 事件信息 事件信息，json，{key:value}结构
func (obj *_IotDeviceEventStatusMgr) WithEventInfoMap(EventInfoMap string) Option {
	return optionFunc(func(o *options) { o.query["event_info_map"] = EventInfoMap })
}

// WithEventLevel event_level获取 事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)
func (obj *_IotDeviceEventStatusMgr) WithEventLevel(EventLevel int) Option {
	return optionFunc(func(o *options) { o.query["event_level"] = EventLevel })
}

// WithEventStatus event_status获取 事件状态 事件状态（告警1，结束0）
func (obj *_IotDeviceEventStatusMgr) WithEventStatus(EventStatus string) Option {
	return optionFunc(func(o *options) { o.query["event_status"] = EventStatus })
}

// GetByOption 功能选项模式获取
func (obj *_IotDeviceEventStatusMgr) GetByOption(opts ...Option) (result IotDeviceEventStatus, err error) {
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
func (obj *_IotDeviceEventStatusMgr) GetByOptions(opts ...Option) (results []*IotDeviceEventStatus, err error) {
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
func (obj *_IotDeviceEventStatusMgr) GetFromTime(Time int64) (result IotDeviceEventStatus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time = ?", Time).Find(&result).Error

	return
}

// GetBatchFromTime 批量唯一主键查找 落库时间 时序数据库的主键
func (obj *_IotDeviceEventStatusMgr) GetBatchFromTime(Times []int64) (results []*IotDeviceEventStatus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time IN (?)", Times).Find(&results).Error

	return
}

// GetFromEventTime 通过event_time获取内容 事件发生时间 事件发生时间，如果上报为空，取服务器时间。
func (obj *_IotDeviceEventStatusMgr) GetFromEventTime(EventTime time.Time) (results []*IotDeviceEventStatus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_time = ?", EventTime).Find(&results).Error

	return
}

// GetBatchFromEventTime 批量唯一主键查找 事件发生时间 事件发生时间，如果上报为空，取服务器时间。
func (obj *_IotDeviceEventStatusMgr) GetBatchFromEventTime(EventTimes []time.Time) (results []*IotDeviceEventStatus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_time IN (?)", EventTimes).Find(&results).Error

	return
}

// GetFromDeviceModelEventID 通过device_type_event_id获取内容 设备类型事件定义ID-外键 设备类型事件定义ID-外键
func (obj *_IotDeviceEventStatusMgr) GetFromDeviceModelEventID(DeviceModelEventID int64) (results []*IotDeviceEventStatus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_event_id = ?", DeviceModelEventID).Find(&results).Error

	return
}

// GetBatchFromDeviceModelEventID 批量唯一主键查找 设备类型事件定义ID-外键 设备类型事件定义ID-外键
func (obj *_IotDeviceEventStatusMgr) GetBatchFromDeviceModelEventID(DeviceModelEventIDs []int64) (results []*IotDeviceEventStatus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_event_id IN (?)", DeviceModelEventIDs).Find(&results).Error

	return
}

// GetFromEventInfoMap 通过event_info_map获取内容 事件信息 事件信息，json，{key:value}结构
func (obj *_IotDeviceEventStatusMgr) GetFromEventInfoMap(EventInfoMap string) (results []*IotDeviceEventStatus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_info_map = ?", EventInfoMap).Find(&results).Error

	return
}

// GetBatchFromEventInfoMap 批量唯一主键查找 事件信息 事件信息，json，{key:value}结构
func (obj *_IotDeviceEventStatusMgr) GetBatchFromEventInfoMap(EventInfoMaps []string) (results []*IotDeviceEventStatus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_info_map IN (?)", EventInfoMaps).Find(&results).Error

	return
}

// GetFromEventLevel 通过event_level获取内容 事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)
func (obj *_IotDeviceEventStatusMgr) GetFromEventLevel(EventLevel int) (results []*IotDeviceEventStatus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_level = ?", EventLevel).Find(&results).Error

	return
}

// GetBatchFromEventLevel 批量唯一主键查找 事件级别 事件级别(信息INFO、告警ALARM、故障ERROR)
func (obj *_IotDeviceEventStatusMgr) GetBatchFromEventLevel(EventLevels []int) (results []*IotDeviceEventStatus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_level IN (?)", EventLevels).Find(&results).Error

	return
}

// GetFromEventStatus 通过event_status获取内容 事件状态 事件状态（告警1，结束0）
func (obj *_IotDeviceEventStatusMgr) GetFromEventStatus(EventStatus string) (results []*IotDeviceEventStatus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_status = ?", EventStatus).Find(&results).Error

	return
}

// GetBatchFromEventStatus 批量唯一主键查找 事件状态 事件状态（告警1，结束0）
func (obj *_IotDeviceEventStatusMgr) GetBatchFromEventStatus(EventStatuss []string) (results []*IotDeviceEventStatus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("event_status IN (?)", EventStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_IotDeviceEventStatusMgr) FetchByPrimaryKey(Time int64) (result IotDeviceEventStatus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time = ?", Time).Find(&result).Error

	return
}
