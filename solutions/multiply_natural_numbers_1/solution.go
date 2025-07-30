package multiply_natural_numbers_1

import "fmt"

func Run() {
	fmt.Println("Running multiply natural numbers solution")

	result := multiply(5, 4) // Example call to the multiply function
	fmt.Printf("The result of multiplying 5 and 4 is: %d\n", result)
}

func multiply(a, b int) int {
	if b == 0 {
		return 0
	}

	if b < 0 {
		return -multiply(a, -b)
	}

	return a + multiply(a, b-1)

}
