package main

func Soma(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func SomaTodoOResto(numerosParaSoma ...[]int) []int {
	var somas []int

	for _, numeros := range numerosParaSoma {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			final := numeros[1:]
			somas = append(somas, Soma(final))		
		}
	}

	return somas
}
