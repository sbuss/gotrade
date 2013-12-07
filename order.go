package gotrade

type OrderType uint8

const (
	Buy OrderType = iota
	Sell
)

type Order struct {
	Type OrderType
	Amount uint64
	Value uint64
}

func NewOrder(orderType OrderType, amount, value unit64) *Order {
	order := new(Order)
	order.type = OrderType
	order.amount = amount
	order.value = value
	return order
}
