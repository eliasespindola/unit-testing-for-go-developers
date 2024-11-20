package calculator

import (
	"errors"
	"unit-testing-go-developers/database"
)

type DiscountCalculator struct {
	minimumPurchase int
	repository      database.Repository
}

func NewDiscountCalculator(minimumPurchase int, repository database.Repository) (*DiscountCalculator, error) {
	if minimumPurchase == 0 {
		return &DiscountCalculator{}, errors.New("minimum purchase amount could not be 0")
	}

	return &DiscountCalculator{
		minimumPurchase: minimumPurchase,
		repository:      repository,
	}, nil
}

func (c DiscountCalculator) Calculate(purchaseAmount int) int {
	if purchaseAmount > c.minimumPurchase {
		multiplier := purchaseAmount / c.minimumPurchase
		discount := c.repository.FindCurrentDiscount()
		return purchaseAmount - (discount * multiplier)
	}

	return purchaseAmount
}
