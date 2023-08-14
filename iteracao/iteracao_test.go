package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	resultado := Repetir("a", 10)
	esperado := "aaaaaaaaaa"

	if resultado != esperado {
		t.Errorf("resultado: %s - esperado: %s", resultado, esperado)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 10)
	}
}

func ExampleRepetir() {
	repetidos := Repetir("e", 4)
	fmt.Println(repetidos)
	// output: eeee
}
