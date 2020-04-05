package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _IotDeviceMgr struct {
	*_BaseMgr
}

// IotDeviceMgr open func
func IotDeviceMgr(db *gorm.DB) *_IotDeviceMgr {
	if db == nil {
		panic(fmt.Errorf("IotDeviceMgr need init by db"))
	}
	return &_IotDeviceMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IotDeviceMgr) GetTableName() string {
	return "iot_device"
}

// Get 获取
func (obj *_IotDeviceMgr) Get() (result IotDevice, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IotDeviceMgr) Gets() (results []*IotDevice, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键 分布式唯一ID
func (obj *_IotDeviceMgr) WithID(ID int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithDeviceUId device_uid获取 设备UID
func (obj *_IotDeviceMgr) WithDeviceUId(DeviceUId string) Option {
	return optionFunc(func(o *options) { o.query["device_uid"] = DeviceUId })
}

// WithDeviceName device_name获取 设备名称
func (obj *_IotDeviceMgr) WithDeviceName(DeviceName string) Option {
	return optionFunc(func(o *options) { o.query["device_name"] = DeviceName })
}

// WithDeviceModelID device_type_id获取 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceMgr) WithDeviceModelID(DeviceModelID int64) Option {
	return optionFunc(func(o *options) { o.query["device_type_id"] = DeviceModelID })
}

// WithOnlineStatus online_status获取 联网状态 在线（1）、离线（0）、未激活（-1）
func (obj *_IotDeviceMgr) WithOnlineStatus(OnlineStatus int) Option {
	return optionFunc(func(o *options) { o.query["online_status"] = OnlineStatus })
}

// GetByOption 功能选项模式获取
func (obj *_IotDeviceMgr) GetByOption(opts ...Option) (result IotDevice, err error) {
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
func (obj *_IotDeviceMgr) GetByOptions(opts ...Option) (results []*IotDevice, err error) {
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

// GetFromID 通过id获取内容 主键 分布式唯一ID
func (obj *_IotDeviceMgr) GetFromID(ID int64) (result IotDevice, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键 分布式唯一ID
func (obj *_IotDeviceMgr) GetBatchFromID(IDs []int64) (results []*IotDevice, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromDeviceUId 通过device_uid获取内容 设备UID
func (obj *_IotDeviceMgr) GetFromDeviceUId(DeviceUId string) (results []*IotDevice, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_uid = ?", DeviceUId).Find(&results).Error

	return
}

// GetBatchFromDeviceUId 批量唯一主键查找 设备UID
func (obj *_IotDeviceMgr) GetBatchFromDeviceUId(DeviceUIds []string) (results []*IotDevice, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_uid IN (?)", DeviceUIds).Find(&results).Error

	return
}

// GetFromDeviceName 通过device_name获取内容 设备名称
func (obj *_IotDeviceMgr) GetFromDeviceName(DeviceName string) (results []*IotDevice, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_name = ?", DeviceName).Find(&results).Error

	return
}

// GetBatchFromDeviceName 批量唯一主键查找 设备名称
func (obj *_IotDeviceMgr) GetBatchFromDeviceName(DeviceNames []string) (results []*IotDevice, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_name IN (?)", DeviceNames).Find(&results).Error

	return
}

// GetFromDeviceModelID 通过device_type_id获取内容 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceMgr) GetFromDeviceModelID(DeviceModelID int64) (results []*IotDevice, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id = ?", DeviceModelID).Find(&results).Error

	return
}

// GetBatchFromDeviceModelID 批量唯一主键查找 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceMgr) GetBatchFromDeviceModelID(DeviceModelIDs []int64) (results []*IotDevice, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id IN (?)", DeviceModelIDs).Find(&results).Error

	return
}

// GetFromOnlineStatus 通过online_status获取内容 联网状态 在线（1）、离线（0）、未激活（-1）
func (obj *_IotDeviceMgr) GetFromOnlineStatus(OnlineStatus int) (results []*IotDevice, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("online_status = ?", OnlineStatus).Find(&results).Error

	return
}

// GetBatchFromOnlineStatus 批量唯一主键查找 联网状态 在线（1）、离线（0）、未激活（-1）
func (obj *_IotDeviceMgr) GetBatchFromOnlineStatus(OnlineStatuss []int) (results []*IotDevice, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("online_status IN (?)", OnlineStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_IotDeviceMgr) FetchByPrimaryKey(ID int64) (result IotDevice, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}
