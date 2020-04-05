package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _IotDeviceModelReadAttrMgr struct {
	*_BaseMgr
}

// IotDeviceModelReadAttrMgr open func
func IotDeviceModelReadAttrMgr(db *gorm.DB) *_IotDeviceModelReadAttrMgr {
	if db == nil {
		panic(fmt.Errorf("IotDeviceModelReadAttrMgr need init by db"))
	}
	return &_IotDeviceModelReadAttrMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IotDeviceModelReadAttrMgr) GetTableName() string {
	return "iot_device_type_read_attr"
}

// Get 获取
func (obj *_IotDeviceModelReadAttrMgr) Get() (result IotDeviceModelReadAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IotDeviceModelReadAttrMgr) Gets() (results []*IotDeviceModelReadAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID 主键
func (obj *_IotDeviceModelReadAttrMgr) WithID(ID int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithDeviceModelID device_type_id获取 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceModelReadAttrMgr) WithDeviceModelID(DeviceModelID int64) Option {
	return optionFunc(func(o *options) { o.query["device_type_id"] = DeviceModelID })
}

// WithAttrPKey attr_p_key获取 只读属性字段名 只读属性字段名
func (obj *_IotDeviceModelReadAttrMgr) WithAttrPKey(AttrPKey string) Option {
	return optionFunc(func(o *options) { o.query["attr_p_key"] = AttrPKey })
}

// WithDataType data_type获取 只读属性-数据类型 只读属性-数据类型
func (obj *_IotDeviceModelReadAttrMgr) WithDataType(DataType string) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = DataType })
}

// WithDefalutValue defalut_value获取 只读属性-默认值 只读属性-默认值
func (obj *_IotDeviceModelReadAttrMgr) WithDefalutValue(DefalutValue string) Option {
	return optionFunc(func(o *options) { o.query["defalut_value"] = DefalutValue })
}

// WithUnit unit获取 只读属性-单位 只读属性-单位
func (obj *_IotDeviceModelReadAttrMgr) WithUnit(Unit string) Option {
	return optionFunc(func(o *options) { o.query["unit"] = Unit })
}

// GetByOption 功能选项模式获取
func (obj *_IotDeviceModelReadAttrMgr) GetByOption(opts ...Option) (result IotDeviceModelReadAttr, err error) {
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
func (obj *_IotDeviceModelReadAttrMgr) GetByOptions(opts ...Option) (results []*IotDeviceModelReadAttr, err error) {
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
func (obj *_IotDeviceModelReadAttrMgr) GetFromID(ID int64) (results []*IotDeviceModelReadAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&results).Error

	return
}

// GetBatchFromID 批量唯一主键查找 ID 主键
func (obj *_IotDeviceModelReadAttrMgr) GetBatchFromID(IDs []int64) (results []*IotDeviceModelReadAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromDeviceModelID 通过device_type_id获取内容 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceModelReadAttrMgr) GetFromDeviceModelID(DeviceModelID int64) (results []*IotDeviceModelReadAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id = ?", DeviceModelID).Find(&results).Error

	return
}

// GetBatchFromDeviceModelID 批量唯一主键查找 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceModelReadAttrMgr) GetBatchFromDeviceModelID(DeviceModelIDs []int64) (results []*IotDeviceModelReadAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id IN (?)", DeviceModelIDs).Find(&results).Error

	return
}

// GetFromAttrPKey 通过attr_p_key获取内容 只读属性字段名 只读属性字段名
func (obj *_IotDeviceModelReadAttrMgr) GetFromAttrPKey(AttrPKey string) (results []*IotDeviceModelReadAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("attr_p_key = ?", AttrPKey).Find(&results).Error

	return
}

// GetBatchFromAttrPKey 批量唯一主键查找 只读属性字段名 只读属性字段名
func (obj *_IotDeviceModelReadAttrMgr) GetBatchFromAttrPKey(AttrPKeys []string) (results []*IotDeviceModelReadAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("attr_p_key IN (?)", AttrPKeys).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容 只读属性-数据类型 只读属性-数据类型
func (obj *_IotDeviceModelReadAttrMgr) GetFromDataType(DataType string) (results []*IotDeviceModelReadAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("data_type = ?", DataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量唯一主键查找 只读属性-数据类型 只读属性-数据类型
func (obj *_IotDeviceModelReadAttrMgr) GetBatchFromDataType(DataTypes []string) (results []*IotDeviceModelReadAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("data_type IN (?)", DataTypes).Find(&results).Error

	return
}

// GetFromDefalutValue 通过defalut_value获取内容 只读属性-默认值 只读属性-默认值
func (obj *_IotDeviceModelReadAttrMgr) GetFromDefalutValue(DefalutValue string) (results []*IotDeviceModelReadAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("defalut_value = ?", DefalutValue).Find(&results).Error

	return
}

// GetBatchFromDefalutValue 批量唯一主键查找 只读属性-默认值 只读属性-默认值
func (obj *_IotDeviceModelReadAttrMgr) GetBatchFromDefalutValue(DefalutValues []string) (results []*IotDeviceModelReadAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("defalut_value IN (?)", DefalutValues).Find(&results).Error

	return
}

// GetFromUnit 通过unit获取内容 只读属性-单位 只读属性-单位
func (obj *_IotDeviceModelReadAttrMgr) GetFromUnit(Unit string) (results []*IotDeviceModelReadAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("unit = ?", Unit).Find(&results).Error

	return
}

// GetBatchFromUnit 批量唯一主键查找 只读属性-单位 只读属性-单位
func (obj *_IotDeviceModelReadAttrMgr) GetBatchFromUnit(Units []string) (results []*IotDeviceModelReadAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("unit IN (?)", Units).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
