package main

import "fmt"

// Type assertion
func main() {
	var minhaVar interface{} = "Roberto Martins"

	println(minhaVar.(string))

	res, ok := minhaVar.(int)

	println(res, ok)
	fmt.Printf("O valor de res é %v e o valor de ok é %v\n", res, ok)

	// Sem o ok ele dá panic e para
	// a execução do programa caso
	// o tipo não seja o esperado
}
