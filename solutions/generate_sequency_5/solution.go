package generate_sequency_5

import "fmt"

func Run() {
	fmt.Println("Running generate sequence solution")

	result := generate(5)
	fmt.Println("The result of generating the 5th number in the sequence is:", result)
}

func generate(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return n
	}

	return 2*generate(n-1) + 3*generate(n-2)
}
