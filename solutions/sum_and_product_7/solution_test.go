package sum_and_product_7

// Disclaimer: This tests are implemented by AI for educational purposes.

import (
	"testing"
)

func TestSumAndProduct_BaseCases(t *testing.T) {
	sum, prod := sum_and_product(0)
	if sum != 0 || prod != 1 {
		t.Errorf("sum_and_product(0) = (%d, %d), want (0, 1)", sum, prod)
	}

	sum, prod = sum_and_product(-5)
	if sum != 0 || prod != 1 {
		t.Errorf("sum_and_product(-5) = (%d, %d), want (0, 1)", sum, prod)
	}
}

func TestSumAndProduct_SmallNumbers(t *testing.T) {
	tests := []struct {
		n     int
		wantS int
		wantP int
	}{
		{1, 1, 1},
		{2, 3, 2},
		{3, 6, 6},
		{4, 10, 24},
		{5, 15, 120},
	}

	for _, tt := range tests {
		sum, prod := sum_and_product(tt.n)
		if sum != tt.wantS || prod != tt.wantP {
			t.Errorf("sum_and_product(%d) = (%d, %d), want (%d, %d)", tt.n, sum, prod, tt.wantS, tt.wantP)
		}
	}
}

func TestSumAndProduct_LargerNumber(t *testing.T) {
	sum, prod := sum_and_product(10)
	if sum != 55 || prod != 3628800 {
		t.Errorf("sum_and_product(10) = (%d, %d), want (55, 3628800)", sum, prod)
	}
}
