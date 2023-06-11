package main

import "testing"

func TestOla(t *testing.T) {
	checkCorrectMessage := func (t *testing.T, result, expected string)  {
		t.Helper()

		if result != expected {
			t.Errorf("result: '%s', expected: '%s'", result, expected)
		}
	}

	t.Run("should say hello to people", func(t *testing.T) {
		result := Hello("Eric", "")
		expected := "Olá, Eric"
	
		checkCorrectMessage(t, result, expected)
	})

	t.Run("should say 'Hello world' when an empty string is passed", func(t *testing.T) {
		result := Hello("", "")
		expected := "Olá, Mundo"

		checkCorrectMessage(t, result, expected)
	})

	t.Run("should say hello in Spanish", func(t *testing.T) {
		result := Hello("Eric", "spanish")
		expected := "Hola, Eric"

		checkCorrectMessage(t, result, expected)
	})

	t.Run("should say hello in French", func(t *testing.T) {
		result := Hello("Eric", "french")
		expected := "Bonjour, Eric"

		checkCorrectMessage(t, result, expected)
	})
}