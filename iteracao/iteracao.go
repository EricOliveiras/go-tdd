package iteracao

import "strings"

// Repetir, repete um número de vezes um caractere
func Repetir(caractere string, quantidadesDeRepeticoes int) string {
	return strings.Repeat(caractere, quantidadesDeRepeticoes)
}
