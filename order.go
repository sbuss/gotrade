package gotrade

type OrderType uint8

const (
	Buy OrderType = iota
	Sell
)

type Order struct {
	Type   OrderType
	Amount uint64
	Value  uint64
}

func NewOrder(orderType OrderType, amount, value uint64) *Order {
	order := new(Order)
	order.Type = orderType
	order.Amount = amount
	order.Value = value
	return order
}
