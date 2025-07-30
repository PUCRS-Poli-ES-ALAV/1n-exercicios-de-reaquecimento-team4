package ackerman_sequence4

import "fmt"

func Run(){
	fmt.Println("Ackerman Sequence")

	ackermanValues := ackermanSequence(3, 4)

	fmt.Printf("Ackerman Sequence(3, 4) = %d\n", ackermanValues)
}

func ackermanSequence(m, n int) int{
	if m == 0 {
		return n + 1
	}

	if m != 0 && n == 0 {
		return ackermanSequence(m-1, 1)
	}

	return ackermanSequence(m-1, ackermanSequence(m, n-1))
}