package sum_and_product_7

import "fmt"

func Run() {
	fmt.Println("Running sum product solution")

	sum, product := sum_and_product(5)
	fmt.Println("The result of sum for 5 is:", sum)
	fmt.Println("The result of product for 5 is:", product)
}

func sum_and_product(n int) (int /* sum */, int /* product */) {
	if n <= 0 {
		return 0, 1
	}

	sum, product := sum_and_product(n - 1)
	sum += n
	product *= n

	return sum, product
}
