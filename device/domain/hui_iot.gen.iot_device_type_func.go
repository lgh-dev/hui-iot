package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _IotDeviceTypeFuncMgr struct {
	*_BaseMgr
}

// IotDeviceTypeFuncMgr open func
func IotDeviceTypeFuncMgr(db *gorm.DB) *_IotDeviceTypeFuncMgr {
	if db == nil {
		panic(fmt.Errorf("IotDeviceTypeFuncMgr need init by db"))
	}
	return &_IotDeviceTypeFuncMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IotDeviceTypeFuncMgr) GetTableName() string {
	return "iot_device_type_func"
}

// Get 获取
func (obj *_IotDeviceTypeFuncMgr) Get() (result IotDeviceTypeFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IotDeviceTypeFuncMgr) Gets() (results []*IotDeviceTypeFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID 主键
func (obj *_IotDeviceTypeFuncMgr) WithID(ID int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithDeviceTypeID device_type_id获取 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceTypeFuncMgr) WithDeviceTypeID(DeviceTypeID int64) Option {
	return optionFunc(func(o *options) { o.query["device_type_id"] = DeviceTypeID })
}

// WithFuncName func_name获取 功能名称 功能名称
func (obj *_IotDeviceTypeFuncMgr) WithFuncName(FuncName string) Option {
	return optionFunc(func(o *options) { o.query["func_name"] = FuncName })
}

// WithInPara in_para获取 输入参数信息 输入参数信息[key:value:data_type,key:value:data_type]
func (obj *_IotDeviceTypeFuncMgr) WithInPara(InPara string) Option {
	return optionFunc(func(o *options) { o.query["in_para"] = InPara })
}

// GetByOption 功能选项模式获取
func (obj *_IotDeviceTypeFuncMgr) GetByOption(opts ...Option) (result IotDeviceTypeFunc, err error) {
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
func (obj *_IotDeviceTypeFuncMgr) GetByOptions(opts ...Option) (results []*IotDeviceTypeFunc, err error) {
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
func (obj *_IotDeviceTypeFuncMgr) GetFromID(ID int64) (results []*IotDeviceTypeFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&results).Error

	return
}

// GetBatchFromID 批量唯一主键查找 ID 主键
func (obj *_IotDeviceTypeFuncMgr) GetBatchFromID(IDs []int64) (results []*IotDeviceTypeFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromDeviceTypeID 通过device_type_id获取内容 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceTypeFuncMgr) GetFromDeviceTypeID(DeviceTypeID int64) (results []*IotDeviceTypeFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id = ?", DeviceTypeID).Find(&results).Error

	return
}

// GetBatchFromDeviceTypeID 批量唯一主键查找 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceTypeFuncMgr) GetBatchFromDeviceTypeID(DeviceTypeIDs []int64) (results []*IotDeviceTypeFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id IN (?)", DeviceTypeIDs).Find(&results).Error

	return
}

// GetFromFuncName 通过func_name获取内容 功能名称 功能名称
func (obj *_IotDeviceTypeFuncMgr) GetFromFuncName(FuncName string) (results []*IotDeviceTypeFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("func_name = ?", FuncName).Find(&results).Error

	return
}

// GetBatchFromFuncName 批量唯一主键查找 功能名称 功能名称
func (obj *_IotDeviceTypeFuncMgr) GetBatchFromFuncName(FuncNames []string) (results []*IotDeviceTypeFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("func_name IN (?)", FuncNames).Find(&results).Error

	return
}

// GetFromInPara 通过in_para获取内容 输入参数信息 输入参数信息[key:value:data_type,key:value:data_type]
func (obj *_IotDeviceTypeFuncMgr) GetFromInPara(InPara string) (results []*IotDeviceTypeFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("in_para = ?", InPara).Find(&results).Error

	return
}

// GetBatchFromInPara 批量唯一主键查找 输入参数信息 输入参数信息[key:value:data_type,key:value:data_type]
func (obj *_IotDeviceTypeFuncMgr) GetBatchFromInPara(InParas []string) (results []*IotDeviceTypeFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("in_para IN (?)", InParas).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
