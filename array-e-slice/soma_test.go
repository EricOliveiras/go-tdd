package arrayeslice

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {
	verificaResultado := func(t *testing.T, resultado, esperado int, dado []int) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado: %d - esperado: %d - dado: %v", resultado, esperado, dado)
		}
	}

	t.Run("coleção de 5 números", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}
	
		resultado := Soma(numeros)
		esperado := 15
	
		verificaResultado(t, resultado, esperado, numeros)
	})

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}
	
		resultado := Soma(numeros)
		esperado := 6
	
		verificaResultado(t, resultado, esperado, numeros)
	})
}

func TestSomaTudo(t *testing.T) {
	resultado := SomaTudo([]int{1,2}, []int{6,3})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("resultado: %v - esperado: %v", resultado, esperado)
	}
}

func TestSomaTodoOResto(t *testing.T) {
	verificaSomas := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado: %v - esperado: %v", resultado, esperado)
		}
	}

	t.Run("faz a soma de alguns slices", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{1,2}, []int{6,9})
		esperado := []int{2, 9}
	
		verificaSomas(t, resultado, esperado)
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{}, []int{6,9})
		esperado := []int{0, 9}
	
		verificaSomas(t, resultado, esperado)
	})
}
