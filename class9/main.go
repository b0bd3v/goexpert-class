package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(10, 11, 12, 12, 13))
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
