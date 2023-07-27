package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(10, 10)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(valor)
}

// func sum(x, y int) int {
// 	return x + y
// }

// func sum(x, y int) (int, bool) {
// 	if x+y >= 50 {
// 		return x + y, true
// 	}

// 	return x + y, false
// }

func sum(x, y int) (int, error) {
	if x+y >= 50 {
		return 0, errors.New("Valor muito alto")
	}

	return x + y, nil
}
