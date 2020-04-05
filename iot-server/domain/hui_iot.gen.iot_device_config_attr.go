package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _IotDeviceConfigAttrMgr struct {
	*_BaseMgr
}

// IotDeviceConfigAttrMgr open func
func IotDeviceConfigAttrMgr(db *gorm.DB) *_IotDeviceConfigAttrMgr {
	if db == nil {
		panic(fmt.Errorf("IotDeviceConfigAttrMgr need init by db"))
	}
	return &_IotDeviceConfigAttrMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IotDeviceConfigAttrMgr) GetTableName() string {
	return "iot_device_config_attr"
}

// Get 获取
func (obj *_IotDeviceConfigAttrMgr) Get() (result IotDeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IotDeviceConfigAttrMgr) Gets() (results []*IotDeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID 主键
func (obj *_IotDeviceConfigAttrMgr) WithID(ID int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithDeviceID device_id获取 设备ID-外键 设备ID-外键
func (obj *_IotDeviceConfigAttrMgr) WithDeviceID(DeviceID int64) Option {
	return optionFunc(func(o *options) { o.query["device_id"] = DeviceID })
}

// WithDeviceModelConfigAttrID device_type_config_attr_id获取 配置属性ID 配置属性ID
func (obj *_IotDeviceConfigAttrMgr) WithDeviceModelConfigAttrID(DeviceModelConfigAttrID string) Option {
	return optionFunc(func(o *options) { o.query["device_type_config_attr_id"] = DeviceModelConfigAttrID })
}

// WithValue value获取 配置属性-值 配置属性-值
func (obj *_IotDeviceConfigAttrMgr) WithValue(Value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = Value })
}

// GetByOption 功能选项模式获取
func (obj *_IotDeviceConfigAttrMgr) GetByOption(opts ...Option) (result IotDeviceConfigAttr, err error) {
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
func (obj *_IotDeviceConfigAttrMgr) GetByOptions(opts ...Option) (results []*IotDeviceConfigAttr, err error) {
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
func (obj *_IotDeviceConfigAttrMgr) GetFromID(ID int64) (result IotDeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 ID 主键
func (obj *_IotDeviceConfigAttrMgr) GetBatchFromID(IDs []int64) (results []*IotDeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromDeviceID 通过device_id获取内容 设备ID-外键 设备ID-外键
func (obj *_IotDeviceConfigAttrMgr) GetFromDeviceID(DeviceID int64) (results []*IotDeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_id = ?", DeviceID).Find(&results).Error

	return
}

// GetBatchFromDeviceID 批量唯一主键查找 设备ID-外键 设备ID-外键
func (obj *_IotDeviceConfigAttrMgr) GetBatchFromDeviceID(DeviceIDs []int64) (results []*IotDeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_id IN (?)", DeviceIDs).Find(&results).Error

	return
}

// GetFromDeviceModelConfigAttrID 通过device_type_config_attr_id获取内容 配置属性ID 配置属性ID
func (obj *_IotDeviceConfigAttrMgr) GetFromDeviceModelConfigAttrID(DeviceModelConfigAttrID string) (results []*IotDeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_config_attr_id = ?", DeviceModelConfigAttrID).Find(&results).Error

	return
}

// GetBatchFromDeviceModelConfigAttrID 批量唯一主键查找 配置属性ID 配置属性ID
func (obj *_IotDeviceConfigAttrMgr) GetBatchFromDeviceModelConfigAttrID(DeviceModelConfigAttrIDs []string) (results []*IotDeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_config_attr_id IN (?)", DeviceModelConfigAttrIDs).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容 配置属性-值 配置属性-值
func (obj *_IotDeviceConfigAttrMgr) GetFromValue(Value string) (results []*IotDeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("value = ?", Value).Find(&results).Error

	return
}

// GetBatchFromValue 批量唯一主键查找 配置属性-值 配置属性-值
func (obj *_IotDeviceConfigAttrMgr) GetBatchFromValue(Values []string) (results []*IotDeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("value IN (?)", Values).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_IotDeviceConfigAttrMgr) FetchByPrimaryKey(ID int64) (result IotDeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}
