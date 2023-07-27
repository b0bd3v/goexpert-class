package main

type ID int

var (
	b bool
	c int
	d string
	e float64
	f ID = 123
)

func main() {
	a := "Hello, World"
	println(a, b, c, d, e, f)
}
