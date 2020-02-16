package sevice

import (
	"hui-iot/device/domain"
	. "hui-iot/device/domain/dto"
)

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 22:18
 * @Desc: Service's interface of device
 */
// Single service interface
type BaseSingleAddService interface {
	add(be domain.BaseEntity) error
}
type BaseSingleUpdateService interface {
	update(be domain.BaseEntity) error
}
type BaseSingleQueryService interface {
	query(id EntityID) (domain.BaseEntity, error)
}
type BaseSingleDeleteService interface {
	Delete(id EntityID) error
}

// batch service interface
type BaseBatchAddService interface {
	adds(be []domain.BaseEntity) error
}
type BaseBatchUpdateService interface {
	updates(be []domain.BaseEntity) error
}
type BaseBatchQueryService interface {
	querys(queryDTO BaseQueryDTO) ([]domain.BaseEntity, error)
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

type DeviceBatchServce interface {
	BaseBatchAddService
	BaseBatchUpdateService
	BaseBatchQueryService
	BaseBatchDeleteService
}

type DeviceSingleServceImpl struct {
}
