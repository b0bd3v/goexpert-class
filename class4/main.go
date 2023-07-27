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
	fmt.Printf("O tipo de E é %T", f)
}
