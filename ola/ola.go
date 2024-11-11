package main

import "fmt"

const espanhol = "espanhol"
const frances = "francês"
const prefixoOlaPortuges = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixoDeSaudacao(idioma) + nome
}

func prefixoDeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	default:
		prefixo = prefixoOlaPortuges
	}
	return
}

func main() {
	fmt.Printf(Ola("mundo", ""))
}
