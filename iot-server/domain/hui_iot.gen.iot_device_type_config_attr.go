package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _IotDeviceTypeConfigAttrMgr struct {
	*_BaseMgr
}

// IotDeviceTypeConfigAttrMgr open func
func IotDeviceTypeConfigAttrMgr(db *gorm.DB) *_IotDeviceTypeConfigAttrMgr {
	if db == nil {
		panic(fmt.Errorf("IotDeviceTypeConfigAttrMgr need init by db"))
	}
	return &_IotDeviceTypeConfigAttrMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IotDeviceTypeConfigAttrMgr) GetTableName() string {
	return "iot_device_type_config_attr"
}

// Get 获取
func (obj *_IotDeviceTypeConfigAttrMgr) Get() (result IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IotDeviceTypeConfigAttrMgr) Gets() (results []*IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID 主键
func (obj *_IotDeviceTypeConfigAttrMgr) WithID(ID int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithDeviceTypeID device_type_id获取 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceTypeConfigAttrMgr) WithDeviceTypeID(DeviceTypeID int64) Option {
	return optionFunc(func(o *options) { o.query["device_type_id"] = DeviceTypeID })
}

// WithAttrPKey attr_p_key获取 配置属性字段名 配置属性字段名
func (obj *_IotDeviceTypeConfigAttrMgr) WithAttrPKey(AttrPKey string) Option {
	return optionFunc(func(o *options) { o.query["attr_p_key"] = AttrPKey })
}

// WithDataType data_type获取 配置属性-数据类型 配置属性-数据类型
func (obj *_IotDeviceTypeConfigAttrMgr) WithDataType(DataType string) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = DataType })
}

// WithDefalutValue defalut_value获取 配置属性-默认值 配置属性-默认值
func (obj *_IotDeviceTypeConfigAttrMgr) WithDefalutValue(DefalutValue string) Option {
	return optionFunc(func(o *options) { o.query["defalut_value"] = DefalutValue })
}

// WithUnit unit获取 配置属性-单位 配置属性-单位
func (obj *_IotDeviceTypeConfigAttrMgr) WithUnit(Unit string) Option {
	return optionFunc(func(o *options) { o.query["unit"] = Unit })
}

// WithMaxValue max_value获取 最大值 最大值
func (obj *_IotDeviceTypeConfigAttrMgr) WithMaxValue(MaxValue string) Option {
	return optionFunc(func(o *options) { o.query["max_value"] = MaxValue })
}

// WithMinValue min_value获取 最小值 最小值
func (obj *_IotDeviceTypeConfigAttrMgr) WithMinValue(MinValue string) Option {
	return optionFunc(func(o *options) { o.query["min_value"] = MinValue })
}

// GetByOption 功能选项模式获取
func (obj *_IotDeviceTypeConfigAttrMgr) GetByOption(opts ...Option) (result IotDeviceTypeConfigAttr, err error) {
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
func (obj *_IotDeviceTypeConfigAttrMgr) GetByOptions(opts ...Option) (results []*IotDeviceTypeConfigAttr, err error) {
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
func (obj *_IotDeviceTypeConfigAttrMgr) GetFromID(ID int64) (results []*IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&results).Error

	return
}

// GetBatchFromID 批量唯一主键查找 ID 主键
func (obj *_IotDeviceTypeConfigAttrMgr) GetBatchFromID(IDs []int64) (results []*IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromDeviceTypeID 通过device_type_id获取内容 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceTypeConfigAttrMgr) GetFromDeviceTypeID(DeviceTypeID int64) (results []*IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id = ?", DeviceTypeID).Find(&results).Error

	return
}

// GetBatchFromDeviceTypeID 批量唯一主键查找 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceTypeConfigAttrMgr) GetBatchFromDeviceTypeID(DeviceTypeIDs []int64) (results []*IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id IN (?)", DeviceTypeIDs).Find(&results).Error

	return
}

// GetFromAttrPKey 通过attr_p_key获取内容 配置属性字段名 配置属性字段名
func (obj *_IotDeviceTypeConfigAttrMgr) GetFromAttrPKey(AttrPKey string) (results []*IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("attr_p_key = ?", AttrPKey).Find(&results).Error

	return
}

// GetBatchFromAttrPKey 批量唯一主键查找 配置属性字段名 配置属性字段名
func (obj *_IotDeviceTypeConfigAttrMgr) GetBatchFromAttrPKey(AttrPKeys []string) (results []*IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("attr_p_key IN (?)", AttrPKeys).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容 配置属性-数据类型 配置属性-数据类型
func (obj *_IotDeviceTypeConfigAttrMgr) GetFromDataType(DataType string) (results []*IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("data_type = ?", DataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量唯一主键查找 配置属性-数据类型 配置属性-数据类型
func (obj *_IotDeviceTypeConfigAttrMgr) GetBatchFromDataType(DataTypes []string) (results []*IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("data_type IN (?)", DataTypes).Find(&results).Error

	return
}

// GetFromDefalutValue 通过defalut_value获取内容 配置属性-默认值 配置属性-默认值
func (obj *_IotDeviceTypeConfigAttrMgr) GetFromDefalutValue(DefalutValue string) (results []*IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("defalut_value = ?", DefalutValue).Find(&results).Error

	return
}

// GetBatchFromDefalutValue 批量唯一主键查找 配置属性-默认值 配置属性-默认值
func (obj *_IotDeviceTypeConfigAttrMgr) GetBatchFromDefalutValue(DefalutValues []string) (results []*IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("defalut_value IN (?)", DefalutValues).Find(&results).Error

	return
}

// GetFromUnit 通过unit获取内容 配置属性-单位 配置属性-单位
func (obj *_IotDeviceTypeConfigAttrMgr) GetFromUnit(Unit string) (results []*IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("unit = ?", Unit).Find(&results).Error

	return
}

// GetBatchFromUnit 批量唯一主键查找 配置属性-单位 配置属性-单位
func (obj *_IotDeviceTypeConfigAttrMgr) GetBatchFromUnit(Units []string) (results []*IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("unit IN (?)", Units).Find(&results).Error

	return
}

// GetFromMaxValue 通过max_value获取内容 最大值 最大值
func (obj *_IotDeviceTypeConfigAttrMgr) GetFromMaxValue(MaxValue string) (results []*IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("max_value = ?", MaxValue).Find(&results).Error

	return
}

// GetBatchFromMaxValue 批量唯一主键查找 最大值 最大值
func (obj *_IotDeviceTypeConfigAttrMgr) GetBatchFromMaxValue(MaxValues []string) (results []*IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("max_value IN (?)", MaxValues).Find(&results).Error

	return
}

// GetFromMinValue 通过min_value获取内容 最小值 最小值
func (obj *_IotDeviceTypeConfigAttrMgr) GetFromMinValue(MinValue string) (results []*IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("min_value = ?", MinValue).Find(&results).Error

	return
}

// GetBatchFromMinValue 批量唯一主键查找 最小值 最小值
func (obj *_IotDeviceTypeConfigAttrMgr) GetBatchFromMinValue(MinValues []string) (results []*IotDeviceTypeConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("min_value IN (?)", MinValues).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
