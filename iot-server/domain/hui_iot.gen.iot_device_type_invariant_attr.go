package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _IotDeviceModelInvariantAttrMgr struct {
	*_BaseMgr
}

// IotDeviceModelInvariantAttrMgr open func
func IotDeviceModelInvariantAttrMgr(db *gorm.DB) *_IotDeviceModelInvariantAttrMgr {
	if db == nil {
		panic(fmt.Errorf("IotDeviceModelInvariantAttrMgr need init by db"))
	}
	return &_IotDeviceModelInvariantAttrMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IotDeviceModelInvariantAttrMgr) GetTableName() string {
	return "iot_device_type_invariant_attr"
}

// Get 获取
func (obj *_IotDeviceModelInvariantAttrMgr) Get() (result IotDeviceModelInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IotDeviceModelInvariantAttrMgr) Gets() (results []*IotDeviceModelInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID 主键
func (obj *_IotDeviceModelInvariantAttrMgr) WithID(ID int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithDeviceModelID device_type_id获取 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceModelInvariantAttrMgr) WithDeviceModelID(DeviceModelID int64) Option {
	return optionFunc(func(o *options) { o.query["device_type_id"] = DeviceModelID })
}

// WithAttrPKey attr_p_key获取 固定属性-字段名 固定属性-字段名
func (obj *_IotDeviceModelInvariantAttrMgr) WithAttrPKey(AttrPKey string) Option {
	return optionFunc(func(o *options) { o.query["attr_p_key"] = AttrPKey })
}

// WithDataType data_type获取 固定属性-数据类型 固定属性-数据类型
func (obj *_IotDeviceModelInvariantAttrMgr) WithDataType(DataType string) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = DataType })
}

// WithDefalutValue defalut_value获取 固定属性-默认值 固定属性-默认值
func (obj *_IotDeviceModelInvariantAttrMgr) WithDefalutValue(DefalutValue string) Option {
	return optionFunc(func(o *options) { o.query["defalut_value"] = DefalutValue })
}

// WithUnit unit获取 固定属性-单位 固定属性-单位
func (obj *_IotDeviceModelInvariantAttrMgr) WithUnit(Unit string) Option {
	return optionFunc(func(o *options) { o.query["unit"] = Unit })
}

// GetByOption 功能选项模式获取
func (obj *_IotDeviceModelInvariantAttrMgr) GetByOption(opts ...Option) (result IotDeviceModelInvariantAttr, err error) {
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
func (obj *_IotDeviceModelInvariantAttrMgr) GetByOptions(opts ...Option) (results []*IotDeviceModelInvariantAttr, err error) {
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
func (obj *_IotDeviceModelInvariantAttrMgr) GetFromID(ID int64) (results []*IotDeviceModelInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&results).Error

	return
}

// GetBatchFromID 批量唯一主键查找 ID 主键
func (obj *_IotDeviceModelInvariantAttrMgr) GetBatchFromID(IDs []int64) (results []*IotDeviceModelInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromDeviceModelID 通过device_type_id获取内容 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceModelInvariantAttrMgr) GetFromDeviceModelID(DeviceModelID int64) (results []*IotDeviceModelInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id = ?", DeviceModelID).Find(&results).Error

	return
}

// GetBatchFromDeviceModelID 批量唯一主键查找 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceModelInvariantAttrMgr) GetBatchFromDeviceModelID(DeviceModelIDs []int64) (results []*IotDeviceModelInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id IN (?)", DeviceModelIDs).Find(&results).Error

	return
}

// GetFromAttrPKey 通过attr_p_key获取内容 固定属性-字段名 固定属性-字段名
func (obj *_IotDeviceModelInvariantAttrMgr) GetFromAttrPKey(AttrPKey string) (results []*IotDeviceModelInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("attr_p_key = ?", AttrPKey).Find(&results).Error

	return
}

// GetBatchFromAttrPKey 批量唯一主键查找 固定属性-字段名 固定属性-字段名
func (obj *_IotDeviceModelInvariantAttrMgr) GetBatchFromAttrPKey(AttrPKeys []string) (results []*IotDeviceModelInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("attr_p_key IN (?)", AttrPKeys).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容 固定属性-数据类型 固定属性-数据类型
func (obj *_IotDeviceModelInvariantAttrMgr) GetFromDataType(DataType string) (results []*IotDeviceModelInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("data_type = ?", DataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量唯一主键查找 固定属性-数据类型 固定属性-数据类型
func (obj *_IotDeviceModelInvariantAttrMgr) GetBatchFromDataType(DataTypes []string) (results []*IotDeviceModelInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("data_type IN (?)", DataTypes).Find(&results).Error

	return
}

// GetFromDefalutValue 通过defalut_value获取内容 固定属性-默认值 固定属性-默认值
func (obj *_IotDeviceModelInvariantAttrMgr) GetFromDefalutValue(DefalutValue string) (results []*IotDeviceModelInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("defalut_value = ?", DefalutValue).Find(&results).Error

	return
}

// GetBatchFromDefalutValue 批量唯一主键查找 固定属性-默认值 固定属性-默认值
func (obj *_IotDeviceModelInvariantAttrMgr) GetBatchFromDefalutValue(DefalutValues []string) (results []*IotDeviceModelInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("defalut_value IN (?)", DefalutValues).Find(&results).Error

	return
}

// GetFromUnit 通过unit获取内容 固定属性-单位 固定属性-单位
func (obj *_IotDeviceModelInvariantAttrMgr) GetFromUnit(Unit string) (results []*IotDeviceModelInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("unit = ?", Unit).Find(&results).Error

	return
}

// GetBatchFromUnit 批量唯一主键查找 固定属性-单位 固定属性-单位
func (obj *_IotDeviceModelInvariantAttrMgr) GetBatchFromUnit(Units []string) (results []*IotDeviceModelInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("unit IN (?)", Units).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
