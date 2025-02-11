package usecase

//
//import(
//	"github.com/jonecoboy/netCheck/internal/core/domain/entity/order"
//	"github.com/jonecoboy/netCheck/internal/core/domain/repository"
//	ninaRouter "github.com/jonecoboy/nina/router"
//)
//
//type OrderInputDTO struct{
//	ID        string   `json:"id"`
//	Price     float64  `json:"price"`
//	Tax       float64  `json:"tax"`
//}
//
//type OrderOutputDTO struct{
//	ID          string    `json:"id"`
//	Price       float64   `json:"price"`
//	Tax         float64   `json:"tax"`
//	FinalPrice  float64   `json:"final_price"`
//}
//
//// tudo baseado em interfaces! Inversão de controle
//type CreateOrderUseCase struct {
//	OrderRepository   repository.OrderRepositoryInterface
//	OrderCreated      ninaRouter.GenericXML //evento
//	EventDispatcher   events.EventDispatcherInterface //quem dispara o evento
//}
////construtor
//func NewCreateOrderUseCase(
//	entity.OrderRepositoryInterface,
//	events.EventInterface,
//	events.EventDispatcherInterface
//) *CreateOrderUseCase{
//	return &CreateOrderUseCase{
//		OrderRepository: OrderRepository,
//		OrderCreated:    OrderCreated,
//		EventDispatcher: EventDispatcher,
//	}
//}
//
////método principal
//// vai transformar o DTO numa order e executar a ordem de negócio, nesse caso o Calculate Final Price
//func (c *CreateOrderUseCase) Execute(input OrderInputDTO) (OrderOutputDTO, error){
//	order := entity.Order{
//		ID:         input.ID,
//		Price:      input.Price,
//		Tax:        input.Tax,
//	}
//	order.CalculateFinalPrice()
//	if err := c.OrderRepository.Save(&order); err != nil {
//		return OrderOutputDTO{},err
//	}
//	dto := OrderOutputDTO{
//		ID:          order.ID,
//		Price:       order.Price,
//		Tax:         order.Tax,
//		FinalPrice:  order.Price + order.Tax,
//	}
//
//	c.OrderCreated.SetPayload(dto)
//	//vai disparar o evento
//	c.EventDispatcher.Dispatch(c.OrderCreated)
//	return dto, nil
//}
