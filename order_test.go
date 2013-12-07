package gotrade

import (
	"testing"
)

func TestOrderCreation(t *testing.T) {
	var (
		amount       uint64
		value_amount uint64
	)
	amount = 1000
	value_amount = 50
	for _, buyOrSell := range []OrderType{Buy, Sell} {
		order := NewOrder(buyOrSell, amount, value_amount)
		if order.Type != buyOrSell {
			t.Errorf("%v != %v", order.Type, buyOrSell)
		}
		if order.Amount != amount {
			t.Errorf("%v != %v", order.Amount, amount)
		}
		if order.Value != value_amount {
			t.Errorf("%v != %v", order.Value, value_amount)
		}
	}
}
