package calculator

import "testing"

type DiscountRepositoryMock struct{}

func (d DiscountRepositoryMock) FindCurrentDiscount() int {
	return 20
}

func TestDiscountCalculator_Calculate(t *testing.T) {
	type testCase struct {
		name                  string
		minimumPurchaseAmount int
		discount              int
		purchaseAmount        int
		expectedAmount        int
	}
	testCases := []testCase{
		{name: "should apply 20", minimumPurchaseAmount: 100, purchaseAmount: 150, expectedAmount: 130},
		{name: "should apply 40", minimumPurchaseAmount: 100, purchaseAmount: 200, expectedAmount: 160},
		{name: "should apply 60", minimumPurchaseAmount: 100, purchaseAmount: 350, expectedAmount: 290},
		{name: "should not apply", minimumPurchaseAmount: 100, purchaseAmount: 50, expectedAmount: 50},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			discountRepositoryMock := DiscountRepositoryMock{}
			calculator, _ := NewDiscountCalculator(tc.minimumPurchaseAmount, discountRepositoryMock)
			amount := calculator.Calculate(tc.purchaseAmount)
			if amount != tc.expectedAmount {
				t.Errorf("Expected amount to be %d, but got %d", tc.expectedAmount, amount)
			}
		})
	}
}

func TestDiscountCalculatorShouldFailWithZeroMininumAmount(t *testing.T) {
	_, err := NewDiscountCalculator(0, DiscountRepositoryMock{})
	if err == nil {
		t.Fatalf("should not create the calculator with 0 minimum purchase amount")
	}
}
