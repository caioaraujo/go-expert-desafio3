package usecase

import (
	"github.com/devfullcycle/14-Clean-Arch/internal/entity"
	"github.com/devfullcycle/14-Clean-Arch/pkg/events"
)

type OrderListDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrdersUseCase) Execute() ([]OrderListDTO, error) {

	orders, err := c.OrderRepository.List()
	var dto_list []OrderListDTO

	if err != nil {
		return []OrderListDTO{}, err
	}

	for _, order := range orders {
		dto := OrderListDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		dto_list = append(dto_list, dto)
	}

	c.OrderCreated.SetPayload(dto_list)
	c.EventDispatcher.Dispatch(c.OrderCreated)

	return dto_list, nil
}
