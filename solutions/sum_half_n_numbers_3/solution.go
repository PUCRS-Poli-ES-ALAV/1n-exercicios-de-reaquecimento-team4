package sum_half_n_numbers_3

import "fmt"

func Run() {
	fmt.Println("Running summing half of the first 15 natural numbers solution")

	result := sum(3)
	fmt.Printf("The result of summing half of the first 3 natural numbers is: %f\n", result)
}

func sum(n int) float64 {
	if n <= 0 {
		return 0
	}
	return 1/float64(n) + sum(n-1)
}
