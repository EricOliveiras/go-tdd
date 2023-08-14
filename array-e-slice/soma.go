package arrayeslice

func Soma(numeros []int) int {
	var soma int
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func SomaTudo(numerosParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numerosParaSomar {
		somas = append(somas, Soma(numeros))
	}
	return somas
}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numerosParaSomar {
		if len(numeros) == 0 {
			somas = append(somas, 0)
			continue
		}
		final := numeros[1:]
		somas = append(somas, Soma(final))
	}
	return somas
}
