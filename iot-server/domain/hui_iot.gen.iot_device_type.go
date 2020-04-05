package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _IotDeviceModelMgr struct {
	*_BaseMgr
}

// IotDeviceModelMgr open func
func IotDeviceModelMgr(db *gorm.DB) *_IotDeviceModelMgr {
	if db == nil {
		panic(fmt.Errorf("IotDeviceModelMgr need init by db"))
	}
	return &_IotDeviceModelMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IotDeviceModelMgr) GetTableName() string {
	return "iot_device_type"
}

// Get 获取
func (obj *_IotDeviceModelMgr) Get() (result IotDeviceModel, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IotDeviceModelMgr) Gets() (results []*IotDeviceModel, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键 分布式唯一ID
func (obj *_IotDeviceModelMgr) WithID(ID int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithDeviceModelUId device_type_uid获取 设备类型UID
func (obj *_IotDeviceModelMgr) WithDeviceModelUId(DeviceModelUId string) Option {
	return optionFunc(func(o *options) { o.query["device_type_uid"] = DeviceModelUId })
}

// WithDeviceModelName device_type_name获取 设备类型名称
func (obj *_IotDeviceModelMgr) WithDeviceModelName(DeviceModelName string) Option {
	return optionFunc(func(o *options) { o.query["device_type_name"] = DeviceModelName })
}

// GetByOption 功能选项模式获取
func (obj *_IotDeviceModelMgr) GetByOption(opts ...Option) (result IotDeviceModel, err error) {
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
func (obj *_IotDeviceModelMgr) GetByOptions(opts ...Option) (results []*IotDeviceModel, err error) {
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
func (obj *_IotDeviceModelMgr) GetFromID(ID int64) (result IotDeviceModel, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键 分布式唯一ID
func (obj *_IotDeviceModelMgr) GetBatchFromID(IDs []int64) (results []*IotDeviceModel, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromDeviceModelUId 通过device_type_uid获取内容 设备类型UID
func (obj *_IotDeviceModelMgr) GetFromDeviceModelUId(DeviceModelUId string) (results []*IotDeviceModel, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_uid = ?", DeviceModelUId).Find(&results).Error

	return
}

// GetBatchFromDeviceModelUId 批量唯一主键查找 设备类型UID
func (obj *_IotDeviceModelMgr) GetBatchFromDeviceModelUId(DeviceModelUIds []string) (results []*IotDeviceModel, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_uid IN (?)", DeviceModelUIds).Find(&results).Error

	return
}

// GetFromDeviceModelName 通过device_type_name获取内容 设备类型名称
func (obj *_IotDeviceModelMgr) GetFromDeviceModelName(DeviceModelName string) (results []*IotDeviceModel, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_name = ?", DeviceModelName).Find(&results).Error

	return
}

// GetBatchFromDeviceModelName 批量唯一主键查找 设备类型名称
func (obj *_IotDeviceModelMgr) GetBatchFromDeviceModelName(DeviceModelNames []string) (results []*IotDeviceModel, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_name IN (?)", DeviceModelNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_IotDeviceModelMgr) FetchByPrimaryKey(ID int64) (result IotDeviceModel, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}
