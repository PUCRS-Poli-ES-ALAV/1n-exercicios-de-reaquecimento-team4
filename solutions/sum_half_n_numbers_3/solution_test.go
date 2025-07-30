package sum_half_n_numbers_3

// Disclaimer: This tests are implemented by AI for educational purposes.

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		n        int
		expected float64
	}{
		{0, 0},
		{1, 1},
		{2, 1.5},          // 1 + 1/2 = 1.5
		{3, 1.8333333333}, // 1 + 1/2 + 1/3 ≈ 1.8333333333
		{4, 2.0833333333}, // 1 + 1/2 + 1/3 + 1/4 ≈ 2.0833333333
		{10, 2.928968254}, // 1 + 1/2 + ... + 1/10 ≈ 2.928968254
		{-5, 0},
		{100, 5.1873775176}, // 1 + 1/2 + ... + 1/100 ≈ 5.1873775176
	}

	for _, tt := range tests {
		t.Run(
			"n="+string(rune(tt.n)),
			func(t *testing.T) {
				result := sum(tt.n)
				if diff := result - tt.expected; diff < -1e-9 || diff > 1e-9 {
					t.Errorf("sum(%d) = %f; want %f", tt.n, result, tt.expected)
				}
			},
		)
	}
}
