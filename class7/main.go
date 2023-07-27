package main

import "fmt"

// import "fmt"

func main() {
	salarios := map[string]int{"Roberto": 1000, "Gustavo": 5000, "Rafael": 3000}

	fmt.Println(salarios["Roberto"])
	delete(salarios, "Roberto")
	fmt.Println(salarios["Roberto"])
	salarios["Bob"] = 10
	fmt.Println(salarios["Bob"])

	// sal := make(map[string]int)
	sal1 := map[string]int{}

	sal1["Roberto"] = 1000

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}
}
