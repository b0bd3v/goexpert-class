package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	roberto := Cliente{
		Nome:  "Roberto",
		Idade: 30,
		Ativo: true,
	}

	roberto.Ativo = false

	fmt.Printf("O cliente %s tem %d anos e está ativo? %t", roberto.Nome, roberto.Idade, roberto.Ativo)
}
