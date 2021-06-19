package facade

import (
	"design-patterns-by-golang/02_structural_patterns/10_facade/delivery"
	"design-patterns-by-golang/02_structural_patterns/10_facade/inventory"
	"errors"
	"fmt"
)

type Shopping interface {
	Buy(buyer, item string) bool
}

type OnlineShopping struct{}

func (*OnlineShopping) Buy(buyer, item string) (bool, error) {
	// 1. check item inventory
	if !inventory.TakeOut(item) {
		return false, fmt.Errorf("%s out of stock", item)
	}
	// 2. delivery item to the buyer
	if !delivery.SendExpress(buyer, item) {
		return false, errors.New("destination unreachable")
	}
	return true, nil
}

func NewOnlineShopping() *OnlineShopping {
	return &OnlineShopping{}
}
