package main

import (
	"fmt"

	"curso-go-mod/matematica"
)

func main() {
	soma := matematica.Soma(10, 20)
	fmt.Println("Resultad: ", soma)
}

// Metodos no Go
// primeira letra maiúscula é pública
// primeira letra minúscula é privada
// Serve para structs
// Para atributos da struct também
