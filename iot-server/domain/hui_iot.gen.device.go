package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _DeviceMgr struct {
	*_BaseMgr
}

// DeviceMgr open func
func DeviceMgr(db *gorm.DB) *_DeviceMgr {
	if db == nil {
		panic(fmt.Errorf("DeviceMgr need init by db"))
	}
	return &_DeviceMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DeviceMgr) GetTableName() string {
	return "device"
}

// Get 获取
func (obj *_DeviceMgr) Get() (result Device, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DeviceMgr) Gets() (results []*Device, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键 分布式唯一ID
func (obj *_DeviceMgr) WithID(ID string) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithDeviceUId device_uid获取 设备UID
func (obj *_DeviceMgr) WithDeviceUId(DeviceUId string) Option {
	return optionFunc(func(o *options) { o.query["device_uid"] = DeviceUId })
}

// WithDeviceName device_name获取 设备名称
func (obj *_DeviceMgr) WithDeviceName(DeviceName string) Option {
	return optionFunc(func(o *options) { o.query["device_name"] = DeviceName })
}

// WithGatewayID gateway_id获取 网关ID
func (obj *_DeviceMgr) WithGatewayID(GatewayID string) Option {
	return optionFunc(func(o *options) { o.query["gateway_id"] = GatewayID })
}

// WithDeviceModelID device_model_id获取 设备模型ID-外键 设备模型ID-外键
func (obj *_DeviceMgr) WithDeviceModelID(DeviceModelID string) Option {
	return optionFunc(func(o *options) { o.query["device_model_id"] = DeviceModelID })
}

// WithOnlineStatus online_status获取 联网状态 在线（1）、离线（0）、未激活（-1）
func (obj *_DeviceMgr) WithOnlineStatus(OnlineStatus int) Option {
	return optionFunc(func(o *options) { o.query["online_status"] = OnlineStatus })
}

// GetByOption 功能选项模式获取
func (obj *_DeviceMgr) GetByOption(opts ...Option) (result Device, err error) {
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
func (obj *_DeviceMgr) GetByOptions(opts ...Option) (results []*Device, err error) {
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
func (obj *_DeviceMgr) GetFromID(ID string) (result Device, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键 分布式唯一ID
func (obj *_DeviceMgr) GetBatchFromID(IDs []string) (results []*Device, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromDeviceUId 通过device_uid获取内容 设备UID
func (obj *_DeviceMgr) GetFromDeviceUId(DeviceUId string) (results []*Device, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_uid = ?", DeviceUId).Find(&results).Error

	return
}

// GetBatchFromDeviceUId 批量唯一主键查找 设备UID
func (obj *_DeviceMgr) GetBatchFromDeviceUId(DeviceUIds []string) (results []*Device, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_uid IN (?)", DeviceUIds).Find(&results).Error

	return
}

// GetFromDeviceName 通过device_name获取内容 设备名称
func (obj *_DeviceMgr) GetFromDeviceName(DeviceName string) (results []*Device, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_name = ?", DeviceName).Find(&results).Error

	return
}

// GetBatchFromDeviceName 批量唯一主键查找 设备名称
func (obj *_DeviceMgr) GetBatchFromDeviceName(DeviceNames []string) (results []*Device, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_name IN (?)", DeviceNames).Find(&results).Error

	return
}

// GetFromGatewayID 通过gateway_id获取内容 网关ID
func (obj *_DeviceMgr) GetFromGatewayID(GatewayID string) (results []*Device, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("gateway_id = ?", GatewayID).Find(&results).Error

	return
}

// GetBatchFromGatewayID 批量唯一主键查找 网关ID
func (obj *_DeviceMgr) GetBatchFromGatewayID(GatewayIDs []string) (results []*Device, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("gateway_id IN (?)", GatewayIDs).Find(&results).Error

	return
}

// GetFromDeviceModelID 通过device_model_id获取内容 设备模型ID-外键 设备模型ID-外键
func (obj *_DeviceMgr) GetFromDeviceModelID(DeviceModelID string) (results []*Device, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_model_id = ?", DeviceModelID).Find(&results).Error

	return
}

// GetBatchFromDeviceModelID 批量唯一主键查找 设备模型ID-外键 设备模型ID-外键
func (obj *_DeviceMgr) GetBatchFromDeviceModelID(DeviceModelIDs []string) (results []*Device, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_model_id IN (?)", DeviceModelIDs).Find(&results).Error

	return
}

// GetFromOnlineStatus 通过online_status获取内容 联网状态 在线（1）、离线（0）、未激活（-1）
func (obj *_DeviceMgr) GetFromOnlineStatus(OnlineStatus int) (results []*Device, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("online_status = ?", OnlineStatus).Find(&results).Error

	return
}

// GetBatchFromOnlineStatus 批量唯一主键查找 联网状态 在线（1）、离线（0）、未激活（-1）
func (obj *_DeviceMgr) GetBatchFromOnlineStatus(OnlineStatuss []int) (results []*Device, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("online_status IN (?)", OnlineStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_DeviceMgr) FetchByPrimaryKey(ID string) (result Device, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}
