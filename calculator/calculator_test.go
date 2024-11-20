package calculator

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type DiscountRepositoryMock struct {
	mock.Mock
}

func (d DiscountRepositoryMock) FindCurrentDiscount() int {
	called := d.Called()
	return called.Int(0)
}

func TestDiscountCalculator_Calculate(t *testing.T) {
	type testCase struct {
		name                  string
		minimumPurchaseAmount int
		purchaseAmount        int
		discount              int
		expectedAmount        int
	}
	testCases := []testCase{
		{
			name:                  "should apply 20",
			minimumPurchaseAmount: 100,
			purchaseAmount:        150,
			discount:              20,
			expectedAmount:        130,
		},
		{
			name:                  "should apply 40",
			minimumPurchaseAmount: 100,
			purchaseAmount:        200,
			discount:              20,
			expectedAmount:        160,
		},
		{
			name:                  "should apply 60",
			minimumPurchaseAmount: 100,
			purchaseAmount:        350,
			discount:              20,
			expectedAmount:        290,
		},
		{
			name:                  "should not apply",
			minimumPurchaseAmount: 100,
			purchaseAmount:        50,
			discount:              20,
			expectedAmount:        50,
		},
		{
			name:                  "should not apply when discount is zero",
			minimumPurchaseAmount: 100,
			purchaseAmount:        50,
			discount:              0,
			expectedAmount:        50,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			discountRepositoryMock := DiscountRepositoryMock{}
			discountRepositoryMock.On("FindCurrentDiscount").Return(tc.discount)
			calculator, _ := NewDiscountCalculator(tc.minimumPurchaseAmount, discountRepositoryMock)
			amount := calculator.Calculate(tc.purchaseAmount)
			assert.Equal(t, tc.expectedAmount, amount)
		})
	}
}

func TestDiscountCalculatorShouldFailWithZeroMininumAmount(t *testing.T) {
	_, err := NewDiscountCalculator(0, DiscountRepositoryMock{})
	if err == nil {
		t.Fatalf("should not create the calculator with 0 minimum purchase amount")
	}
}
