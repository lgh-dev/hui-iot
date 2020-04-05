package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _IotDeviceModelFuncMgr struct {
	*_BaseMgr
}

// IotDeviceModelFuncMgr open func
func IotDeviceModelFuncMgr(db *gorm.DB) *_IotDeviceModelFuncMgr {
	if db == nil {
		panic(fmt.Errorf("IotDeviceModelFuncMgr need init by db"))
	}
	return &_IotDeviceModelFuncMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_IotDeviceModelFuncMgr) GetTableName() string {
	return "iot_device_type_func"
}

// Get 获取
func (obj *_IotDeviceModelFuncMgr) Get() (result IotDeviceModelFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_IotDeviceModelFuncMgr) Gets() (results []*IotDeviceModelFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID 主键
func (obj *_IotDeviceModelFuncMgr) WithID(ID int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithDeviceModelID device_type_id获取 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceModelFuncMgr) WithDeviceModelID(DeviceModelID int64) Option {
	return optionFunc(func(o *options) { o.query["device_type_id"] = DeviceModelID })
}

// WithFuncName func_name获取 功能名称 功能名称
func (obj *_IotDeviceModelFuncMgr) WithFuncName(FuncName string) Option {
	return optionFunc(func(o *options) { o.query["func_name"] = FuncName })
}

// WithInPara in_para获取 输入参数信息 输入参数信息[key:value:data_type,key:value:data_type]
func (obj *_IotDeviceModelFuncMgr) WithInPara(InPara string) Option {
	return optionFunc(func(o *options) { o.query["in_para"] = InPara })
}

// GetByOption 功能选项模式获取
func (obj *_IotDeviceModelFuncMgr) GetByOption(opts ...Option) (result IotDeviceModelFunc, err error) {
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
func (obj *_IotDeviceModelFuncMgr) GetByOptions(opts ...Option) (results []*IotDeviceModelFunc, err error) {
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
func (obj *_IotDeviceModelFuncMgr) GetFromID(ID int64) (results []*IotDeviceModelFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&results).Error

	return
}

// GetBatchFromID 批量唯一主键查找 ID 主键
func (obj *_IotDeviceModelFuncMgr) GetBatchFromID(IDs []int64) (results []*IotDeviceModelFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromDeviceModelID 通过device_type_id获取内容 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceModelFuncMgr) GetFromDeviceModelID(DeviceModelID int64) (results []*IotDeviceModelFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id = ?", DeviceModelID).Find(&results).Error

	return
}

// GetBatchFromDeviceModelID 批量唯一主键查找 设备类型ID-外键 设备类型ID-外键
func (obj *_IotDeviceModelFuncMgr) GetBatchFromDeviceModelID(DeviceModelIDs []int64) (results []*IotDeviceModelFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_type_id IN (?)", DeviceModelIDs).Find(&results).Error

	return
}

// GetFromFuncName 通过func_name获取内容 功能名称 功能名称
func (obj *_IotDeviceModelFuncMgr) GetFromFuncName(FuncName string) (results []*IotDeviceModelFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("func_name = ?", FuncName).Find(&results).Error

	return
}

// GetBatchFromFuncName 批量唯一主键查找 功能名称 功能名称
func (obj *_IotDeviceModelFuncMgr) GetBatchFromFuncName(FuncNames []string) (results []*IotDeviceModelFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("func_name IN (?)", FuncNames).Find(&results).Error

	return
}

// GetFromInPara 通过in_para获取内容 输入参数信息 输入参数信息[key:value:data_type,key:value:data_type]
func (obj *_IotDeviceModelFuncMgr) GetFromInPara(InPara string) (results []*IotDeviceModelFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("in_para = ?", InPara).Find(&results).Error

	return
}

// GetBatchFromInPara 批量唯一主键查找 输入参数信息 输入参数信息[key:value:data_type,key:value:data_type]
func (obj *_IotDeviceModelFuncMgr) GetBatchFromInPara(InParas []string) (results []*IotDeviceModelFunc, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("in_para IN (?)", InParas).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
