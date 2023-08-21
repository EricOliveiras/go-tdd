package main

import (
	"bytes"
	"testing"
)

func TestCumprimenta(t *testing.T) {
	buffer := bytes.Buffer{}
	Cumprimenta(&buffer, "Eric")

	resultado := buffer.String()
	esperado := "Olá, Eric"

	if resultado != esperado {
		t.Errorf("resultado: %s - esperado: %s", resultado, esperado)
	}
}
