package main

import "fmt"

type ID int

var (
	a string = "Hello, World"
	b bool
	c int
	d string
	e float64
	f ID = 123
)

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	// fmt.Printf("O tipo de E é %T", f)
	// fmt.Println(meuArray[len(meuArray)-2])

	for i, v := range meuArray {
		fmt.Printf("O valor de i é %d e o valor de v é %d\n", i, v)
	}
}
