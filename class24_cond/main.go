package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	a := 1
	// b := 2
	// c := 3

	// Só tem else if

	// if a > b {
	// 	fmt.Println("A é maior que B")
	// } else {
	// 	fmt.Println("B é maior que A")
	// }

	// if a > b && c > a {
	// 	fmt.Println("A é maior que B e C é maior que A")
	// }

	switch a {
	case 1:
		fmt.Println("A é igual a 1")
	case 2:
		fmt.Println("A é igual a 2")
	case 3:
		fmt.Println("A é igual a 3")
	default:
		fmt.Println("A não é igual a 1, 2 ou 3")
	}

}
