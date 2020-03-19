package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _IotDeviceTypeMgr struct {
	*_BaseMgr
}

// IotDeviceTypeMgr open func
func IotDeviceTypeMgr(db *gorm.DB) *_IotDeviceTypeMgr {
	if db == nil {
		panic(fmt.Errorf("IotDeviceTypeMgr need init by db"))
	}
	return &_IotDeviceTypeMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IotDeviceTypeMgr) GetTableName() string {
	return "iot_device_type"
}

// Get 获取
func (obj *_IotDeviceTypeMgr) Get() (result IotDeviceType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IotDeviceTypeMgr) Gets() (results []*IotDeviceType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键 分布式唯一ID
func (obj *_IotDeviceTypeMgr) WithID(ID int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithDeviceTypeUId device_type_uid获取 设备类型UID
func (obj *_IotDeviceTypeMgr) WithDeviceTypeUId(DeviceTypeUId string) Option {
	return optionFunc(func(o *options) { o.query["device_type_uid"] = DeviceTypeUId })
}

// WithDeviceTypeName device_type_name获取 设备类型名称
func (obj *_IotDeviceTypeMgr) WithDeviceTypeName(DeviceTypeName string) Option {
	return optionFunc(func(o *options) { o.query["device_type_name"] = DeviceTypeName })
}

// GetByOption 功能选项模式获取
func (obj *_IotDeviceTypeMgr) GetByOption(opts ...Option) (result IotDeviceType, err error) {
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
func (obj *_IotDeviceTypeMgr) GetByOptions(opts ...Option) (results []*IotDeviceType, err error) {
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
func (obj *_IotDeviceTypeMgr) GetFromID(ID int64) (result IotDeviceType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键 分布式唯一ID
func (obj *_IotDeviceTypeMgr) GetBatchFromID(IDs []int64) (results []*IotDeviceType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromDeviceTypeUId 通过device_type_uid获取内容 设备类型UID
func (obj *_IotDeviceTypeMgr) GetFromDeviceTypeUId(DeviceTypeUId string) (results []*IotDeviceType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_uid = ?", DeviceTypeUId).Find(&results).Error

	return
}

// GetBatchFromDeviceTypeUId 批量唯一主键查找 设备类型UID
func (obj *_IotDeviceTypeMgr) GetBatchFromDeviceTypeUId(DeviceTypeUIds []string) (results []*IotDeviceType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_uid IN (?)", DeviceTypeUIds).Find(&results).Error

	return
}

// GetFromDeviceTypeName 通过device_type_name获取内容 设备类型名称
func (obj *_IotDeviceTypeMgr) GetFromDeviceTypeName(DeviceTypeName string) (results []*IotDeviceType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_name = ?", DeviceTypeName).Find(&results).Error

	return
}

// GetBatchFromDeviceTypeName 批量唯一主键查找 设备类型名称
func (obj *_IotDeviceTypeMgr) GetBatchFromDeviceTypeName(DeviceTypeNames []string) (results []*IotDeviceType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_name IN (?)", DeviceTypeNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_IotDeviceTypeMgr) FetchByPrimaryKey(ID int64) (result IotDeviceType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}
