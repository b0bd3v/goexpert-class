package main

import "fmt"


func main() {
	fmt.Println("Hello World"))	
}

// Go build 
// gera no so local

// GOOS=windows go build main.go
// GOOS=linux go build main.go
// GOOS=darwin go build main.go
// GOOS=freebsd go build main.go
// GOOS=netbsd go build main.go
// GOOS=machos go build main.go

// por padrao a variavel de ambiente GOOS é darwin


// GOOS=darwin GOARCH=amd64 go build main.go

// Para listar os sistemas operacionais e arquiteturas
// go tool dist list

// go env GOOS GOARCH

// go build vai gerar um executavel

// Parametro -o para definir o nome do executavel

