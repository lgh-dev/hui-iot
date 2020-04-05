package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _IotDeviceModelConfigAttrMgr struct {
	*_BaseMgr
}

// IotDeviceModelConfigAttrMgr open func
func IotDeviceModelConfigAttrMgr(db *gorm.DB) *_IotDeviceModelConfigAttrMgr {
	if db == nil {
		panic(fmt.Errorf("IotDeviceModelConfigAttrMgr need init by db"))
	}
	return &_IotDeviceModelConfigAttrMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IotDeviceModelConfigAttrMgr) GetTableName() string {
	return "iot_device_type_config_attr"
}

// Get 获取
func (obj *_IotDeviceModelConfigAttrMgr) Get() (result IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IotDeviceModelConfigAttrMgr) Gets() (results []*IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID 主键
func (obj *_IotDeviceModelConfigAttrMgr) WithID(ID int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithDeviceModelID device_type_id获取 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceModelConfigAttrMgr) WithDeviceModelID(DeviceModelID int64) Option {
	return optionFunc(func(o *options) { o.query["device_type_id"] = DeviceModelID })
}

// WithAttrPKey attr_p_key获取 配置属性字段名 配置属性字段名
func (obj *_IotDeviceModelConfigAttrMgr) WithAttrPKey(AttrPKey string) Option {
	return optionFunc(func(o *options) { o.query["attr_p_key"] = AttrPKey })
}

// WithDataType data_type获取 配置属性-数据类型 配置属性-数据类型
func (obj *_IotDeviceModelConfigAttrMgr) WithDataType(DataType string) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = DataType })
}

// WithDefalutValue defalut_value获取 配置属性-默认值 配置属性-默认值
func (obj *_IotDeviceModelConfigAttrMgr) WithDefalutValue(DefalutValue string) Option {
	return optionFunc(func(o *options) { o.query["defalut_value"] = DefalutValue })
}

// WithUnit unit获取 配置属性-单位 配置属性-单位
func (obj *_IotDeviceModelConfigAttrMgr) WithUnit(Unit string) Option {
	return optionFunc(func(o *options) { o.query["unit"] = Unit })
}

// WithMaxValue max_value获取 最大值 最大值
func (obj *_IotDeviceModelConfigAttrMgr) WithMaxValue(MaxValue string) Option {
	return optionFunc(func(o *options) { o.query["max_value"] = MaxValue })
}

// WithMinValue min_value获取 最小值 最小值
func (obj *_IotDeviceModelConfigAttrMgr) WithMinValue(MinValue string) Option {
	return optionFunc(func(o *options) { o.query["min_value"] = MinValue })
}

// GetByOption 功能选项模式获取
func (obj *_IotDeviceModelConfigAttrMgr) GetByOption(opts ...Option) (result IotDeviceModelConfigAttr, err error) {
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
func (obj *_IotDeviceModelConfigAttrMgr) GetByOptions(opts ...Option) (results []*IotDeviceModelConfigAttr, err error) {
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
func (obj *_IotDeviceModelConfigAttrMgr) GetFromID(ID int64) (results []*IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&results).Error

	return
}

// GetBatchFromID 批量唯一主键查找 ID 主键
func (obj *_IotDeviceModelConfigAttrMgr) GetBatchFromID(IDs []int64) (results []*IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromDeviceModelID 通过device_type_id获取内容 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceModelConfigAttrMgr) GetFromDeviceModelID(DeviceModelID int64) (results []*IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id = ?", DeviceModelID).Find(&results).Error

	return
}

// GetBatchFromDeviceModelID 批量唯一主键查找 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceModelConfigAttrMgr) GetBatchFromDeviceModelID(DeviceModelIDs []int64) (results []*IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id IN (?)", DeviceModelIDs).Find(&results).Error

	return
}

// GetFromAttrPKey 通过attr_p_key获取内容 配置属性字段名 配置属性字段名
func (obj *_IotDeviceModelConfigAttrMgr) GetFromAttrPKey(AttrPKey string) (results []*IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("attr_p_key = ?", AttrPKey).Find(&results).Error

	return
}

// GetBatchFromAttrPKey 批量唯一主键查找 配置属性字段名 配置属性字段名
func (obj *_IotDeviceModelConfigAttrMgr) GetBatchFromAttrPKey(AttrPKeys []string) (results []*IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("attr_p_key IN (?)", AttrPKeys).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容 配置属性-数据类型 配置属性-数据类型
func (obj *_IotDeviceModelConfigAttrMgr) GetFromDataType(DataType string) (results []*IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("data_type = ?", DataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量唯一主键查找 配置属性-数据类型 配置属性-数据类型
func (obj *_IotDeviceModelConfigAttrMgr) GetBatchFromDataType(DataTypes []string) (results []*IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("data_type IN (?)", DataTypes).Find(&results).Error

	return
}

// GetFromDefalutValue 通过defalut_value获取内容 配置属性-默认值 配置属性-默认值
func (obj *_IotDeviceModelConfigAttrMgr) GetFromDefalutValue(DefalutValue string) (results []*IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("defalut_value = ?", DefalutValue).Find(&results).Error

	return
}

// GetBatchFromDefalutValue 批量唯一主键查找 配置属性-默认值 配置属性-默认值
func (obj *_IotDeviceModelConfigAttrMgr) GetBatchFromDefalutValue(DefalutValues []string) (results []*IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("defalut_value IN (?)", DefalutValues).Find(&results).Error

	return
}

// GetFromUnit 通过unit获取内容 配置属性-单位 配置属性-单位
func (obj *_IotDeviceModelConfigAttrMgr) GetFromUnit(Unit string) (results []*IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("unit = ?", Unit).Find(&results).Error

	return
}

// GetBatchFromUnit 批量唯一主键查找 配置属性-单位 配置属性-单位
func (obj *_IotDeviceModelConfigAttrMgr) GetBatchFromUnit(Units []string) (results []*IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("unit IN (?)", Units).Find(&results).Error

	return
}

// GetFromMaxValue 通过max_value获取内容 最大值 最大值
func (obj *_IotDeviceModelConfigAttrMgr) GetFromMaxValue(MaxValue string) (results []*IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("max_value = ?", MaxValue).Find(&results).Error

	return
}

// GetBatchFromMaxValue 批量唯一主键查找 最大值 最大值
func (obj *_IotDeviceModelConfigAttrMgr) GetBatchFromMaxValue(MaxValues []string) (results []*IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("max_value IN (?)", MaxValues).Find(&results).Error

	return
}

// GetFromMinValue 通过min_value获取内容 最小值 最小值
func (obj *_IotDeviceModelConfigAttrMgr) GetFromMinValue(MinValue string) (results []*IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("min_value = ?", MinValue).Find(&results).Error

	return
}

// GetBatchFromMinValue 批量唯一主键查找 最小值 最小值
func (obj *_IotDeviceModelConfigAttrMgr) GetBatchFromMinValue(MinValues []string) (results []*IotDeviceModelConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("min_value IN (?)", MinValues).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
