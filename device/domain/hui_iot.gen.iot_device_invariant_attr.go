package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _IotDeviceInvariantAttrMgr struct {
	*_BaseMgr
}

// IotDeviceInvariantAttrMgr open func
func IotDeviceInvariantAttrMgr(db *gorm.DB) *_IotDeviceInvariantAttrMgr {
	if db == nil {
		panic(fmt.Errorf("IotDeviceInvariantAttrMgr need init by db"))
	}
	return &_IotDeviceInvariantAttrMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IotDeviceInvariantAttrMgr) GetTableName() string {
	return "iot_device_invariant_attr"
}

// Get 获取
func (obj *_IotDeviceInvariantAttrMgr) Get() (result IotDeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IotDeviceInvariantAttrMgr) Gets() (results []*IotDeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID 主键
func (obj *_IotDeviceInvariantAttrMgr) WithID(ID int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithDeviceID device_id获取 设备ID-外键 设备ID-外键
func (obj *_IotDeviceInvariantAttrMgr) WithDeviceID(DeviceID int64) Option {
	return optionFunc(func(o *options) { o.query["device_id"] = DeviceID })
}

// WithDeviceTypeInvariantAttrID device_type_invariant_attr_id获取 固定属性ID 固定属性ID
func (obj *_IotDeviceInvariantAttrMgr) WithDeviceTypeInvariantAttrID(DeviceTypeInvariantAttrID string) Option {
	return optionFunc(func(o *options) { o.query["device_type_invariant_attr_id"] = DeviceTypeInvariantAttrID })
}

// WithValue value获取 固定属性-值 固定属性-值
func (obj *_IotDeviceInvariantAttrMgr) WithValue(Value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = Value })
}

// GetByOption 功能选项模式获取
func (obj *_IotDeviceInvariantAttrMgr) GetByOption(opts ...Option) (result IotDeviceInvariantAttr, err error) {
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
func (obj *_IotDeviceInvariantAttrMgr) GetByOptions(opts ...Option) (results []*IotDeviceInvariantAttr, err error) {
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
func (obj *_IotDeviceInvariantAttrMgr) GetFromID(ID int64) (result IotDeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 ID 主键
func (obj *_IotDeviceInvariantAttrMgr) GetBatchFromID(IDs []int64) (results []*IotDeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromDeviceID 通过device_id获取内容 设备ID-外键 设备ID-外键
func (obj *_IotDeviceInvariantAttrMgr) GetFromDeviceID(DeviceID int64) (results []*IotDeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_id = ?", DeviceID).Find(&results).Error

	return
}

// GetBatchFromDeviceID 批量唯一主键查找 设备ID-外键 设备ID-外键
func (obj *_IotDeviceInvariantAttrMgr) GetBatchFromDeviceID(DeviceIDs []int64) (results []*IotDeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_id IN (?)", DeviceIDs).Find(&results).Error

	return
}

// GetFromDeviceTypeInvariantAttrID 通过device_type_invariant_attr_id获取内容 固定属性ID 固定属性ID
func (obj *_IotDeviceInvariantAttrMgr) GetFromDeviceTypeInvariantAttrID(DeviceTypeInvariantAttrID string) (results []*IotDeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_invariant_attr_id = ?", DeviceTypeInvariantAttrID).Find(&results).Error

	return
}

// GetBatchFromDeviceTypeInvariantAttrID 批量唯一主键查找 固定属性ID 固定属性ID
func (obj *_IotDeviceInvariantAttrMgr) GetBatchFromDeviceTypeInvariantAttrID(DeviceTypeInvariantAttrIDs []string) (results []*IotDeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_invariant_attr_id IN (?)", DeviceTypeInvariantAttrIDs).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容 固定属性-值 固定属性-值
func (obj *_IotDeviceInvariantAttrMgr) GetFromValue(Value string) (results []*IotDeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("value = ?", Value).Find(&results).Error

	return
}

// GetBatchFromValue 批量唯一主键查找 固定属性-值 固定属性-值
func (obj *_IotDeviceInvariantAttrMgr) GetBatchFromValue(Values []string) (results []*IotDeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("value IN (?)", Values).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_IotDeviceInvariantAttrMgr) FetchByPrimaryKey(ID int64) (result IotDeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}
