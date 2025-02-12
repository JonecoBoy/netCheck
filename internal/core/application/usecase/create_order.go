package usecase

import (
	"github.com/jonecoboy/netCheck/internal/core/domain/entity"
	"github.com/jonecoboy/netCheck/internal/core/domain/repository"
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
	"github.com/jonecoboy/netCheck/pkg/nina"
	"time"
)

type OrderInputDTO struct {
	ID                pkgEntity.ID   `json:"id" bson:"id"`
	UserId            pkgEntity.ID   `json:"user_id" bson:"user_id"`
	BillingAddressId  pkgEntity.ID   `json:"billing_address_id" bson:"billing_address_id"`
	ShippingAddressId pkgEntity.ID   `json:"shipping_address_id" bson:"shipping_address_id"`
	ItemIds           []pkgEntity.ID `json:"item_ids" bson:"item_ids"`
	Price             int            `json:"price" bson:"price"`
	discountCodes     []pkgEntity.ID `json:"discount_codes" bson:"discount_codes"`
}

type OrderOutputDTO struct {
	ID        string             `json:"id" bson:"id"`
	User      entity.User        `json:"user" bson:"user"`
	Addresses entity.Addresses   `json:"addresses" bson:"addresses"`
	Items     []entity.OrderItem `json:"items" bson:"items"`
	Totals    entity.Totals      `json:"final_price"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedAt time.Time          `bson:"deleted_at" json:"deleted_at"`
}
type CreateOrderUseCase struct {
	OrderRepository repository.OrderRepositoryInterface
	OrderCreated    nina.EventInterface // Event
	EventDispatcher nina.EventDispatcherInterface
}

// //construtor
func NewCreateOrderUseCase(
	orderRepository *repository.OrderRepositoryInterface,
	event *nina.EventInterface,
	dispatcher *nina.EventDispatcherInterface,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: *orderRepository,
		OrderCreated:    *event,
		EventDispatcher: *dispatcher,
	}
}

// //método principal
// // vai transformar o DTO numa order e executar a ordem de negócio. Buscar no repositorio tudo, add items, add shipping e calculo dos precos
func (c *CreateOrderUseCase) Execute(o OrderInputDTO) (OrderOutputDTO, error) {
	order := entity.Order{
		ID:    o.ID,
		Price: o.Price,
	}
	order.CalculateFinalPrice()
	if err := c.OrderRepository.Save(&order); err != nil {
		return OrderOutputDTO{}, err
	}
	dto := OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.Price + order.Tax,
	}

	c.OrderCreated.SetPayload(dto)
	//vai disparar o evento
	c.EventDispatcher.Dispatch(c.OrderCreated)
	return dto, nil
}
