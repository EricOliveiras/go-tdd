package main

import (
	"testing"
)

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado: %s - esperado: %s", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Eric", "")
		esperado := "Olá, Eric"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Olá, mundo' qunado uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, Mundo"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz olá em espanhol", func(t *testing.T) {
		resultado := Ola("Eric", "espanhol")
		esperado := "Hola, Eric"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz olá em francês", func(t *testing.T) {
		resultado := Ola("Eric", "frances")
		esperado := "Bonjour, Eric"
		
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
