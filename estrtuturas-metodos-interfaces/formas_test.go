package estrtuturasmetodosinterfaces

import "testing"

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado: %.2f - esperado: %.2f", resultado, esperado)
	}
}

func TestArea(t *testing.T) {
	// verificaArea := func(t *testing.T, forma Forma, esperado float64) {
	// 	t.Helper()

	// 	resultado := forma.Area()

	// 	if resultado != esperado {
	// 		t.Errorf("resultado: %.2f - esperado: %.2f", forma, esperado)
	// 	}
	// }

	testesArea := []struct {
		forma    Forma
		esperado float64
	}{
		{forma: Retangulo{Largura: 12, Altura: 6}, esperado: 72},
		{forma: Circulo{10}, esperado: 314.1592653589793},
		{forma: Triangulo{Base: 12, Altura: 6}, esperado: 36},
	}

	for _, tt := range testesArea {
		resultado := tt.forma.Area()
		if resultado != tt.esperado {
			t.Errorf("%#v: resultado %.2f, esperado %.2f",tt.forma, resultado, tt.esperado)
		}
	}
}
