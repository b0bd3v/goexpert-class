package main

//  Generics

// func somaInteiro(m map[string]int) int {
// 	var soma int
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

// func somaFloat(m map[string]float64) float64 {
// 	var soma int
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

// func main() {
// 	m := map[string]int{"Roberto": 10000, "Fulano": 20000, "Ciclano": 30000}

// 	mFloat := map[string]float64{"Roberto": 10000.0, "Fulano": 20000.0, "Ciclano": 30000.0}

// 	println(somaInteiro(m))

// 	println(somaFloat(mFloat))
// }

type MyNumber int

// Constraints
// Olhar a doc de go sobre constraints

// Number é uma interface que eu criei
type Number interface {
	~int | ~float64
}

// T é um tipo genérico
func soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// comparable é uma interface que já existe no Go
func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Roberto": 10000, "Fulano": 20000, "Ciclano": 30000}

	mFloat := map[string]float64{"Roberto": 10000.0, "Fulano": 20000.0, "Ciclano": 30000.0}

	m3 := map[string]MyNumber{"Roberto": 10000, "Fulano": 20000, "Ciclano": 30000}

	println(soma(m))

	println(soma(mFloat))

	println(soma(m3))

	println(Compara(10, 10.0))
}
