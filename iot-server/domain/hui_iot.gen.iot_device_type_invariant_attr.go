package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _IotDeviceTypeInvariantAttrMgr struct {
	*_BaseMgr
}

// IotDeviceTypeInvariantAttrMgr open func
func IotDeviceTypeInvariantAttrMgr(db *gorm.DB) *_IotDeviceTypeInvariantAttrMgr {
	if db == nil {
		panic(fmt.Errorf("IotDeviceTypeInvariantAttrMgr need init by db"))
	}
	return &_IotDeviceTypeInvariantAttrMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IotDeviceTypeInvariantAttrMgr) GetTableName() string {
	return "iot_device_type_invariant_attr"
}

// Get 获取
func (obj *_IotDeviceTypeInvariantAttrMgr) Get() (result IotDeviceTypeInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IotDeviceTypeInvariantAttrMgr) Gets() (results []*IotDeviceTypeInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID 主键
func (obj *_IotDeviceTypeInvariantAttrMgr) WithID(ID int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithDeviceTypeID device_type_id获取 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceTypeInvariantAttrMgr) WithDeviceTypeID(DeviceTypeID int64) Option {
	return optionFunc(func(o *options) { o.query["device_type_id"] = DeviceTypeID })
}

// WithAttrPKey attr_p_key获取 固定属性-字段名 固定属性-字段名
func (obj *_IotDeviceTypeInvariantAttrMgr) WithAttrPKey(AttrPKey string) Option {
	return optionFunc(func(o *options) { o.query["attr_p_key"] = AttrPKey })
}

// WithDataType data_type获取 固定属性-数据类型 固定属性-数据类型
func (obj *_IotDeviceTypeInvariantAttrMgr) WithDataType(DataType string) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = DataType })
}

// WithDefalutValue defalut_value获取 固定属性-默认值 固定属性-默认值
func (obj *_IotDeviceTypeInvariantAttrMgr) WithDefalutValue(DefalutValue string) Option {
	return optionFunc(func(o *options) { o.query["defalut_value"] = DefalutValue })
}

// WithUnit unit获取 固定属性-单位 固定属性-单位
func (obj *_IotDeviceTypeInvariantAttrMgr) WithUnit(Unit string) Option {
	return optionFunc(func(o *options) { o.query["unit"] = Unit })
}

// GetByOption 功能选项模式获取
func (obj *_IotDeviceTypeInvariantAttrMgr) GetByOption(opts ...Option) (result IotDeviceTypeInvariantAttr, err error) {
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
func (obj *_IotDeviceTypeInvariantAttrMgr) GetByOptions(opts ...Option) (results []*IotDeviceTypeInvariantAttr, err error) {
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
func (obj *_IotDeviceTypeInvariantAttrMgr) GetFromID(ID int64) (results []*IotDeviceTypeInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&results).Error

	return
}

// GetBatchFromID 批量唯一主键查找 ID 主键
func (obj *_IotDeviceTypeInvariantAttrMgr) GetBatchFromID(IDs []int64) (results []*IotDeviceTypeInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromDeviceTypeID 通过device_type_id获取内容 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceTypeInvariantAttrMgr) GetFromDeviceTypeID(DeviceTypeID int64) (results []*IotDeviceTypeInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id = ?", DeviceTypeID).Find(&results).Error

	return
}

// GetBatchFromDeviceTypeID 批量唯一主键查找 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceTypeInvariantAttrMgr) GetBatchFromDeviceTypeID(DeviceTypeIDs []int64) (results []*IotDeviceTypeInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id IN (?)", DeviceTypeIDs).Find(&results).Error

	return
}

// GetFromAttrPKey 通过attr_p_key获取内容 固定属性-字段名 固定属性-字段名
func (obj *_IotDeviceTypeInvariantAttrMgr) GetFromAttrPKey(AttrPKey string) (results []*IotDeviceTypeInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("attr_p_key = ?", AttrPKey).Find(&results).Error

	return
}

// GetBatchFromAttrPKey 批量唯一主键查找 固定属性-字段名 固定属性-字段名
func (obj *_IotDeviceTypeInvariantAttrMgr) GetBatchFromAttrPKey(AttrPKeys []string) (results []*IotDeviceTypeInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("attr_p_key IN (?)", AttrPKeys).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容 固定属性-数据类型 固定属性-数据类型
func (obj *_IotDeviceTypeInvariantAttrMgr) GetFromDataType(DataType string) (results []*IotDeviceTypeInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("data_type = ?", DataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量唯一主键查找 固定属性-数据类型 固定属性-数据类型
func (obj *_IotDeviceTypeInvariantAttrMgr) GetBatchFromDataType(DataTypes []string) (results []*IotDeviceTypeInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("data_type IN (?)", DataTypes).Find(&results).Error

	return
}

// GetFromDefalutValue 通过defalut_value获取内容 固定属性-默认值 固定属性-默认值
func (obj *_IotDeviceTypeInvariantAttrMgr) GetFromDefalutValue(DefalutValue string) (results []*IotDeviceTypeInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("defalut_value = ?", DefalutValue).Find(&results).Error

	return
}

// GetBatchFromDefalutValue 批量唯一主键查找 固定属性-默认值 固定属性-默认值
func (obj *_IotDeviceTypeInvariantAttrMgr) GetBatchFromDefalutValue(DefalutValues []string) (results []*IotDeviceTypeInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("defalut_value IN (?)", DefalutValues).Find(&results).Error

	return
}

// GetFromUnit 通过unit获取内容 固定属性-单位 固定属性-单位
func (obj *_IotDeviceTypeInvariantAttrMgr) GetFromUnit(Unit string) (results []*IotDeviceTypeInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("unit = ?", Unit).Find(&results).Error

	return
}

// GetBatchFromUnit 批量唯一主键查找 固定属性-单位 固定属性-单位
func (obj *_IotDeviceTypeInvariantAttrMgr) GetBatchFromUnit(Units []string) (results []*IotDeviceTypeInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("unit IN (?)", Units).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
