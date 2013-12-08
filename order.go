package gotrade

import (
	"errors"
	"fmt"
	"github.com/nu7hatch/gouuid"
)

type OrderType uint8

const (
	Buy OrderType = iota
	Sell
)

type Order struct {
	Type   OrderType
	Amount uint64
	Value  uint64
	User   *uuid.UUID
}

// Create a new Order
func NewOrder(orderType OrderType, amount, value uint64, user *uuid.UUID) (*Order, error) {
	if orderType != Buy && orderType != Sell {
		errorMsg := fmt.Sprintf("Invalid OrderType. Valid values are [%v, %v]", Buy, Sell)
		return nil, errors.New(errorMsg)
	}
	return &Order{orderType, amount, value, user}, nil
}
