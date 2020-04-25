package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _DeviceConfigAttrMgr struct {
	*_BaseMgr
}

// DeviceConfigAttrMgr open func
func DeviceConfigAttrMgr(db *gorm.DB) *_DeviceConfigAttrMgr {
	if db == nil {
		panic(fmt.Errorf("DeviceConfigAttrMgr need init by db"))
	}
	return &_DeviceConfigAttrMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DeviceConfigAttrMgr) GetTableName() string {
	return "device_config_attr"
}

// Get 获取
func (obj *_DeviceConfigAttrMgr) Get() (result DeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DeviceConfigAttrMgr) Gets() (results []*DeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID 主键
func (obj *_DeviceConfigAttrMgr) WithID(ID int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithDeviceID device_id获取 设备ID-外键
func (obj *_DeviceConfigAttrMgr) WithDeviceID(DeviceID int64) Option {
	return optionFunc(func(o *options) { o.query["device_id"] = DeviceID })
}

// WithUkey ukey获取 配置属性标识
func (obj *_DeviceConfigAttrMgr) WithUkey(Ukey string) Option {
	return optionFunc(func(o *options) { o.query["ukey"] = Ukey })
}

// WithValue value获取 配置属性-值
func (obj *_DeviceConfigAttrMgr) WithValue(Value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = Value })
}

// GetByOption 功能选项模式获取
func (obj *_DeviceConfigAttrMgr) GetByOption(opts ...Option) (result DeviceConfigAttr, err error) {
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
func (obj *_DeviceConfigAttrMgr) GetByOptions(opts ...Option) (results []*DeviceConfigAttr, err error) {
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
func (obj *_DeviceConfigAttrMgr) GetFromID(ID int64) (result DeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 ID 主键
func (obj *_DeviceConfigAttrMgr) GetBatchFromID(IDs []int64) (results []*DeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromDeviceID 通过device_id获取内容 设备ID-外键
func (obj *_DeviceConfigAttrMgr) GetFromDeviceID(DeviceID int64) (results []*DeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_id = ?", DeviceID).Find(&results).Error

	return
}

// GetBatchFromDeviceID 批量唯一主键查找 设备ID-外键
func (obj *_DeviceConfigAttrMgr) GetBatchFromDeviceID(DeviceIDs []int64) (results []*DeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_id IN (?)", DeviceIDs).Find(&results).Error

	return
}

// GetFromUkey 通过ukey获取内容 配置属性标识
func (obj *_DeviceConfigAttrMgr) GetFromUkey(Ukey string) (results []*DeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ukey = ?", Ukey).Find(&results).Error

	return
}

// GetBatchFromUkey 批量唯一主键查找 配置属性标识
func (obj *_DeviceConfigAttrMgr) GetBatchFromUkey(Ukeys []string) (results []*DeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ukey IN (?)", Ukeys).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容 配置属性-值
func (obj *_DeviceConfigAttrMgr) GetFromValue(Value string) (results []*DeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("value = ?", Value).Find(&results).Error

	return
}

// GetBatchFromValue 批量唯一主键查找 配置属性-值
func (obj *_DeviceConfigAttrMgr) GetBatchFromValue(Values []string) (results []*DeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("value IN (?)", Values).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_DeviceConfigAttrMgr) FetchByPrimaryKey(ID int64) (result DeviceConfigAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}
