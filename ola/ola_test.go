package main

import (
	"testing"
)

func TestOla(t *testing.T) {
	resultado := Ola("Eric")
	esperado := "Olá, Eric!"

	if resultado != esperado {
		t.Errorf("resultado: %s - esperado: %s", resultado, esperado)
	}
}
