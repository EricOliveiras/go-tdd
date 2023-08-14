package inteiros

import (
	"fmt"
	"testing"
)

func TestAdicionador(t *testing.T) {
	resultado := Adiciona(2, 2)
	esperado := 4

	if resultado != esperado {
		t.Errorf("resultado: %d - esperado: %d", resultado, esperado)
	}
}

func ExampleAdiciona() {
	soma := Adiciona(1, 5)
	fmt.Println(soma)
	// output: 6
}
