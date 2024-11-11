package main

import "fmt"

const prefixoOlaPortuges = "Olá, "

func Ola(nome string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return prefixoOlaPortuges + nome
}

func main() {
	fmt.Printf(Ola("mundo"))
}
