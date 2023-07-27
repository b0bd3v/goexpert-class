package main

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 100}
}

func (c *Conta) simular(valor int) int {
	// c.nome = "Roberto Martins"
	// fmt.Printf("O cliente %v andou\n", c.nome)
	c.saldo += valor

	return c.saldo
}

func main() {
	conta := Conta{
		saldo: 100,
	}

	conta.simular(200)

	println(conta.saldo)
}
