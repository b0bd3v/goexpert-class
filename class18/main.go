package main

import "fmt"

// interfaces vazias
func main() {
	var x interface{} = 10
	var y interface{} = "Olá mundo"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("Tipo: %T e valor é %v\n", t, t)
}
