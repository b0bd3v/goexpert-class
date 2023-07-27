package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
	Bairro     string
}

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s está desativado", c.Nome)
}

func main() {
	roberto := Cliente{
		Nome:  "Roberto",
		Idade: 30,
		Ativo: true,
	}

	roberto.Ativo = false
	// roberto.Cidade = "São Paulo"
	// roberto.Endereco.Cidade = "São Paulo"
	roberto.Desativar()
	roberto.Address.Cidade = "São Paulo"

	// fmt.Printf("O cliente %s tem %d anos e está ativo? %t", roberto.Nome, roberto.Idade, roberto.Ativo)
}
