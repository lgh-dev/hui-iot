package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _DeviceCommandMgr struct {
	*_BaseMgr
}

// DeviceCommandMgr open func
func DeviceCommandMgr(db *gorm.DB) *_DeviceCommandMgr {
	if db == nil {
		panic(fmt.Errorf("DeviceCommandMgr need init by db"))
	}
	return &_DeviceCommandMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DeviceCommandMgr) GetTableName() string {
	return "device_command"
}

// Get 获取
func (obj *_DeviceCommandMgr) Get() (result DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DeviceCommandMgr) Gets() (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithCmdID cmd_id获取 指令ID
func (obj *_DeviceCommandMgr) WithCmdID(CmdID string) Option {
	return optionFunc(func(o *options) { o.query["cmd_id"] = CmdID })
}

// WithDeviceID device_id获取 设备ID
func (obj *_DeviceCommandMgr) WithDeviceID(DeviceID string) Option {
	return optionFunc(func(o *options) { o.query["device_id"] = DeviceID })
}

// WithSendData send_data获取 发送数据
func (obj *_DeviceCommandMgr) WithSendData(SendData string) Option {
	return optionFunc(func(o *options) { o.query["send_data"] = SendData })
}

// WithReturnData return_data获取 响应数据（同步或异步）
func (obj *_DeviceCommandMgr) WithReturnData(ReturnData string) Option {
	return optionFunc(func(o *options) { o.query["return_data"] = ReturnData })
}

// WithStatus status获取 指令状态(send【下发】、ack【设备已接收】、succ【执行成功】、error【异常】)
func (obj *_DeviceCommandMgr) WithStatus(Status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = Status })
}

// WithMsg msg获取 执行信息，包括成功或失败的信息
func (obj *_DeviceCommandMgr) WithMsg(Msg string) Option {
	return optionFunc(func(o *options) { o.query["msg"] = Msg })
}

// WithIsRetry is_retry获取 是否需要重发
func (obj *_DeviceCommandMgr) WithIsRetry(IsRetry string) Option {
	return optionFunc(func(o *options) { o.query["is_retry"] = IsRetry })
}

// WithRetryInterval retry_interval获取 重发间隔(s)
func (obj *_DeviceCommandMgr) WithRetryInterval(RetryInterval string) Option {
	return optionFunc(func(o *options) { o.query["retry_interval"] = RetryInterval })
}

// WithRetryNumber retry_number获取 重发次数
func (obj *_DeviceCommandMgr) WithRetryNumber(RetryNumber string) Option {
	return optionFunc(func(o *options) { o.query["retry_number"] = RetryNumber })
}

// WithRetryCompletedNumber retry_completed_number获取 已重发次数
func (obj *_DeviceCommandMgr) WithRetryCompletedNumber(RetryCompletedNumber string) Option {
	return optionFunc(func(o *options) { o.query["retry_completed_number"] = RetryCompletedNumber })
}

// GetByOption 功能选项模式获取
func (obj *_DeviceCommandMgr) GetByOption(opts ...Option) (result DeviceCommand, err error) {
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
func (obj *_DeviceCommandMgr) GetByOptions(opts ...Option) (results []*DeviceCommand, err error) {
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

// GetFromCmdID 通过cmd_id获取内容 指令ID
func (obj *_DeviceCommandMgr) GetFromCmdID(CmdID string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cmd_id = ?", CmdID).Find(&results).Error

	return
}

// GetBatchFromCmdID 批量唯一主键查找 指令ID
func (obj *_DeviceCommandMgr) GetBatchFromCmdID(CmdIDs []string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cmd_id IN (?)", CmdIDs).Find(&results).Error

	return
}

// GetFromDeviceID 通过device_id获取内容 设备ID
func (obj *_DeviceCommandMgr) GetFromDeviceID(DeviceID string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_id = ?", DeviceID).Find(&results).Error

	return
}

// GetBatchFromDeviceID 批量唯一主键查找 设备ID
func (obj *_DeviceCommandMgr) GetBatchFromDeviceID(DeviceIDs []string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("device_id IN (?)", DeviceIDs).Find(&results).Error

	return
}

// GetFromSendData 通过send_data获取内容 发送数据
func (obj *_DeviceCommandMgr) GetFromSendData(SendData string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("send_data = ?", SendData).Find(&results).Error

	return
}

// GetBatchFromSendData 批量唯一主键查找 发送数据
func (obj *_DeviceCommandMgr) GetBatchFromSendData(SendDatas []string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("send_data IN (?)", SendDatas).Find(&results).Error

	return
}

// GetFromReturnData 通过return_data获取内容 响应数据（同步或异步）
func (obj *_DeviceCommandMgr) GetFromReturnData(ReturnData string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("return_data = ?", ReturnData).Find(&results).Error

	return
}

// GetBatchFromReturnData 批量唯一主键查找 响应数据（同步或异步）
func (obj *_DeviceCommandMgr) GetBatchFromReturnData(ReturnDatas []string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("return_data IN (?)", ReturnDatas).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 指令状态(send【下发】、ack【设备已接收】、succ【执行成功】、error【异常】)
func (obj *_DeviceCommandMgr) GetFromStatus(Status string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", Status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 指令状态(send【下发】、ack【设备已接收】、succ【执行成功】、error【异常】)
func (obj *_DeviceCommandMgr) GetBatchFromStatus(Statuss []string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", Statuss).Find(&results).Error

	return
}

// GetFromMsg 通过msg获取内容 执行信息，包括成功或失败的信息
func (obj *_DeviceCommandMgr) GetFromMsg(Msg string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("msg = ?", Msg).Find(&results).Error

	return
}

// GetBatchFromMsg 批量唯一主键查找 执行信息，包括成功或失败的信息
func (obj *_DeviceCommandMgr) GetBatchFromMsg(Msgs []string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("msg IN (?)", Msgs).Find(&results).Error

	return
}

// GetFromIsRetry 通过is_retry获取内容 是否需要重发
func (obj *_DeviceCommandMgr) GetFromIsRetry(IsRetry string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_retry = ?", IsRetry).Find(&results).Error

	return
}

// GetBatchFromIsRetry 批量唯一主键查找 是否需要重发
func (obj *_DeviceCommandMgr) GetBatchFromIsRetry(IsRetrys []string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_retry IN (?)", IsRetrys).Find(&results).Error

	return
}

// GetFromRetryInterval 通过retry_interval获取内容 重发间隔(s)
func (obj *_DeviceCommandMgr) GetFromRetryInterval(RetryInterval string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("retry_interval = ?", RetryInterval).Find(&results).Error

	return
}

// GetBatchFromRetryInterval 批量唯一主键查找 重发间隔(s)
func (obj *_DeviceCommandMgr) GetBatchFromRetryInterval(RetryIntervals []string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("retry_interval IN (?)", RetryIntervals).Find(&results).Error

	return
}

// GetFromRetryNumber 通过retry_number获取内容 重发次数
func (obj *_DeviceCommandMgr) GetFromRetryNumber(RetryNumber string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("retry_number = ?", RetryNumber).Find(&results).Error

	return
}

// GetBatchFromRetryNumber 批量唯一主键查找 重发次数
func (obj *_DeviceCommandMgr) GetBatchFromRetryNumber(RetryNumbers []string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("retry_number IN (?)", RetryNumbers).Find(&results).Error

	return
}

// GetFromRetryCompletedNumber 通过retry_completed_number获取内容 已重发次数
func (obj *_DeviceCommandMgr) GetFromRetryCompletedNumber(RetryCompletedNumber string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("retry_completed_number = ?", RetryCompletedNumber).Find(&results).Error

	return
}

// GetBatchFromRetryCompletedNumber 批量唯一主键查找 已重发次数
func (obj *_DeviceCommandMgr) GetBatchFromRetryCompletedNumber(RetryCompletedNumbers []string) (results []*DeviceCommand, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("retry_completed_number IN (?)", RetryCompletedNumbers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
