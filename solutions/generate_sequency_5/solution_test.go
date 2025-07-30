package generate_sequency_5

// Disclaimer: This tests are implemented by AI for educational purposes.

import "testing"

func TestGenerate(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 2*2 + 3*1},   // 4 + 3 = 7
		{4, 2*7 + 3*2},   // 14 + 6 = 20
		{5, 2*20 + 3*7},  // 40 + 21 = 61
		{6, 2*61 + 3*20}, // 122 + 60 = 182
		{0, 0},           // edge case, not defined by recurrence
		{-1, 0},          // edge case, not defined by recurrence
	}

	for _, tt := range tests {
		t.Run(
			"n="+string(rune(tt.n)),
			func(t *testing.T) {
				result := generate(tt.n)
				if result != tt.expected {
					t.Errorf("generate(%d) = %d; want %d", tt.n, result, tt.expected)
				}
			},
		)
	}
}
