package gotrade

import (
	"testing"
	"github.com/nu7hatch/gouuid"
)

func randomUser() (user *uuid.UUID) {
	user, _ = uuid.NewV4()
	return user
}

func TestOrderCreation(t *testing.T) {
	var (
		amount       uint64
		value_amount uint64
	)
	amount = 1000
	value_amount = 50
	for _, buyOrSell := range []OrderType{Buy, Sell} {
		order, err := NewOrder(buyOrSell, amount, value_amount, randomUser())
		if err != nil {
			t.Error("Error instantiating the Order")
		}
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

func TestBadInitValues(t *testing.T) {
	order, err := NewOrder(OrderType(3), uint64(1000), uint64(100), randomUser())
	if err == nil {
		t.Error("Order should not have been created with %v", OrderType(3))
	}
	if order != nil {
		t.Error("Should have not gotten a valid Order")
	}
}

func TestBuyOrder(t *testing.T) {
	order := NewBuyOrder(uint64(1000), uint64(100), randomUser())
	if order == nil {
		t.Error("Order wasn't created successfully")
	}
	if order.Type != Buy {
		t.Error("Expecting a Buy order, but got %v", order.Type)
	}
}

func TestSellOrder(t *testing.T) {
	order := NewSellOrder(uint64(1000), uint64(100), randomUser())
	if order == nil {
		t.Error("Order wasn't created successfully")
	}
	if order.Type != Sell {
		t.Error("Expecting a Sell order, but got %v", order.Type)
	}
}
