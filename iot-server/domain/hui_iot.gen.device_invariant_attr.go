package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _DeviceInvariantAttrMgr struct {
	*_BaseMgr
}

// DeviceInvariantAttrMgr open func
func DeviceInvariantAttrMgr(db *gorm.DB) *_DeviceInvariantAttrMgr {
	if db == nil {
		panic(fmt.Errorf("DeviceInvariantAttrMgr need init by db"))
	}
	return &_DeviceInvariantAttrMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DeviceInvariantAttrMgr) GetTableName() string {
	return "device_invariant_attr"
}

// Get 获取
func (obj *_DeviceInvariantAttrMgr) Get() (result DeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DeviceInvariantAttrMgr) Gets() (results []*DeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID 主键
func (obj *_DeviceInvariantAttrMgr) WithID(ID int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithDeviceID device_id获取 设备ID-外键 设备ID-外键
func (obj *_DeviceInvariantAttrMgr) WithDeviceID(DeviceID int64) Option {
	return optionFunc(func(o *options) { o.query["device_id"] = DeviceID })
}

// WithUkey ukey获取 固定属性标识
func (obj *_DeviceInvariantAttrMgr) WithUkey(Ukey string) Option {
	return optionFunc(func(o *options) { o.query["ukey"] = Ukey })
}

// WithValue value获取 固定属性-值
func (obj *_DeviceInvariantAttrMgr) WithValue(Value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = Value })
}

// GetByOption 功能选项模式获取
func (obj *_DeviceInvariantAttrMgr) GetByOption(opts ...Option) (result DeviceInvariantAttr, err error) {
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
func (obj *_DeviceInvariantAttrMgr) GetByOptions(opts ...Option) (results []*DeviceInvariantAttr, err error) {
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
func (obj *_DeviceInvariantAttrMgr) GetFromID(ID int64) (result DeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 ID 主键
func (obj *_DeviceInvariantAttrMgr) GetBatchFromID(IDs []int64) (results []*DeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromDeviceID 通过device_id获取内容 设备ID-外键 设备ID-外键
func (obj *_DeviceInvariantAttrMgr) GetFromDeviceID(DeviceID int64) (results []*DeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_id = ?", DeviceID).Find(&results).Error

	return
}

// GetBatchFromDeviceID 批量唯一主键查找 设备ID-外键 设备ID-外键
func (obj *_DeviceInvariantAttrMgr) GetBatchFromDeviceID(DeviceIDs []int64) (results []*DeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_id IN (?)", DeviceIDs).Find(&results).Error

	return
}

// GetFromUkey 通过ukey获取内容 固定属性标识
func (obj *_DeviceInvariantAttrMgr) GetFromUkey(Ukey string) (results []*DeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ukey = ?", Ukey).Find(&results).Error

	return
}

// GetBatchFromUkey 批量唯一主键查找 固定属性标识
func (obj *_DeviceInvariantAttrMgr) GetBatchFromUkey(Ukeys []string) (results []*DeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ukey IN (?)", Ukeys).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容 固定属性-值
func (obj *_DeviceInvariantAttrMgr) GetFromValue(Value string) (results []*DeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("value = ?", Value).Find(&results).Error

	return
}

// GetBatchFromValue 批量唯一主键查找 固定属性-值
func (obj *_DeviceInvariantAttrMgr) GetBatchFromValue(Values []string) (results []*DeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("value IN (?)", Values).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_DeviceInvariantAttrMgr) FetchByPrimaryKey(ID int64) (result DeviceInvariantAttr, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}
