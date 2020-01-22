package dto

import (
	. "light-iot/base/domain"
	"time"
)

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 22:21
 * @Desc: device's dto
 */
type BaseQueryDTO struct {
	beginTime time.Time
	endTime   time.Time
}

//Device Add DTO
type DeviceAddDTO struct {
	Device
}

//Device update DTO
type DeviceUpdateDTO struct {
	Device
}

//Device query DTO
type DeviceQueryDTO struct {
	BaseQueryDTO
	Device
}

//Object ID , be used to findByID or deleteByID
type EntityID int64
