package repository

import (
	"github.com/jonecoboy/netCheck/internal/core/domain/entity/order"
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
)

type OrderRepositoryInterface interface {
	Save(order *entity.Order) error
	FindByID(id pkgEntity.ID) (*entity.Order, error)
}
