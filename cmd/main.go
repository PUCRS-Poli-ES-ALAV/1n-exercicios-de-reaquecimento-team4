package main

import (
	"fmt"

	ackerman_sequence4 "github.com/lucasarieta/1n-exercicios-de-reaquecimento-team4/solutions/ackerman_sequence_4"
	"github.com/lucasarieta/1n-exercicios-de-reaquecimento-team4/solutions/alphabet_sequency_9"
	"github.com/lucasarieta/1n-exercicios-de-reaquecimento-team4/solutions/example"
	"github.com/lucasarieta/1n-exercicios-de-reaquecimento-team4/solutions/generate_sequency_5"
	incremental_numbers_sum2 "github.com/lucasarieta/1n-exercicios-de-reaquecimento-team4/solutions/incremental_numbers_sum_2"
	"github.com/lucasarieta/1n-exercicios-de-reaquecimento-team4/solutions/multiply_natural_numbers_1"
	"github.com/lucasarieta/1n-exercicios-de-reaquecimento-team4/solutions/sum_and_product_7"
	"github.com/lucasarieta/1n-exercicios-de-reaquecimento-team4/solutions/sum_half_n_numbers_3"
)

// Run tests with: go test -v ./solutions/*
// Run main with: go run cmd/main.go

func main() {
	fmt.Println("Hello world!")

	// Example run
	example.Run()

	multiply_natural_numbers_1.Run()
	sum_half_n_numbers_3.Run()
	generate_sequency_5.Run()
	incremental_numbers_sum2.Run()
	sum_and_product_7.Run()
	ackerman_sequence4.Run()
	alphabet_sequency_9.Run()
}
