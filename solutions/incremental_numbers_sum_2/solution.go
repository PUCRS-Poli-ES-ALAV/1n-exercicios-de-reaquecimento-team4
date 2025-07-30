package incremental_numbers_sum2

import (
	"fmt"
)

func Run() {
	fmt.Println("Running sum of incremented number")

	result := calculate(3, 5)
	fmt.Println("Result:", result)
}

func calculate(a, b int) int {
	if a == 0 {
		return 0
	}
	return sumOnes(a) + sumOnes(b)
}

func sumOnes(n int) int {
	if n == 0 {
		return 0
	}
	return 1 + sumOnes(n-1)
}
