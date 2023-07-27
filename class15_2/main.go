package main

import "fmt"

func main() {

	valor := 11

	ponteiro := &valor

	*ponteiro = 12

	fmt.Printf("Valor: %v, Ponteiro: %v\n", valor, ponteiro)
}
