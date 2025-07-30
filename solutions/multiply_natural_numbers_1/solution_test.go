package multiply_natural_numbers_1

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{5, 4, 20},
		{3, 7, 21},
		{0, 10, 0},
		{10, 0, 0},
		{-5, 4, -20},
		{5, -4, -20},
		{-5, -4, 20},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d*%d", tt.a, tt.b), func(t *testing.T) {
			result := multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
			fmt.Printf("multiply(%d, %d) = %d\n", tt.a, tt.b, result)
		})
	}
}
