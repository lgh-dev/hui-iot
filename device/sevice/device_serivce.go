package sevice

import (
	. "hui-iot/base/domain"
	. "hui-iot/device/api/dto"
)

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 22:18
 * @Desc: Service's interface of device
 */
// Single service interface
type BaseSingleAddService interface {
	add(be BaseEntity) error
}
type BaseSingleUpdateService interface {
	update(be BaseEntity) error
}
type BaseSingleQueryService interface {
	query(id EntityID) (BaseEntity, error)
}
type BaseSingleDeleteService interface {
	Delete(id EntityID) error
}

// batch service interface
type BaseBatchAddService interface {
	adds(be []BaseEntity) error
}
type BaseBatchUpdateService interface {
	updates(be []BaseEntity) error
}
type BaseBatchQueryService interface {
	querys(queryDTO BaseQueryDTO) ([]BaseEntity, error)
}
type BaseBatchDeleteService interface {
	Deletes(id []EntityID) error
}

type DeviceSingleServce interface {
	BaseSingleAddService
	BaseSingleUpdateService
	BaseSingleQueryService
	BaseSingleDeleteService
}

type DeviceBatchServce struct {
	BaseBatchAddService
	BaseBatchUpdateService
	BaseBatchQueryService
	BaseBatchDeleteService
}

type DeviceSingleServceImpl struct {
}
