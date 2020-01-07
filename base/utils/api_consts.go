package utils

const (
	//全局上下文
	PATH_IOT = "/iot"
	//设备服务上下文
	PATH_DEVICE = "/device"
	//二级上下文
	PATH_SINGLE = "/single"
	PATH_BATCH  = "/batch"

	PATH_IOT_DEVICE_SINGLE = PATH_IOT + PATH_DEVICE + PATH_SINGLE
	PATH_IOT_DEVICE_BATCH  = PATH_IOT + PATH_DEVICE + PATH_BATCH

	//增删查改
	PATH_ADD    = "/add"
	PATH_UPDATE = "/update"
	PATH_DELETE = "/delete"
	PATH_QUERY  = "/query"
	PATH_EXEC   = "/exec"
)
