package main

import "fmt"

func Ola(nome string) string {
	return "Olá, " + nome
}

func main() {
	fmt.Printf(Ola("mundo"))
}

