package gotrade

import (
	"testing"
)

func TestOrderCreation(t *testing.T) {
	buy_amount := 1000
	value_amount := 50
	order = NewOrder(Buy, buy_amount, value_amount)
	if order.Type != Buy {
		t.Errorf("%v != %v", order.Type, Buy)
	}
}
