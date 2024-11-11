package iteracao

func Repetir(caracter string, qtdRepeticoes int) string {
	var repeticoes string
	for i := 0; i < qtdRepeticoes; i++ {
		repeticoes += caracter
	}
	return repeticoes
}
