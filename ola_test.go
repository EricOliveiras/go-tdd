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
		result := Ola("Eric")
		expected := "Olá, Eric"
	
		checkCorrectMessage(t, result, expected)
	})

	t.Run("should say 'Hello world' when an empty string is passed", func(t *testing.T) {
		result := Ola("")
		expected := "Olá, Mundo"

		checkCorrectMessage(t, result, expected)
	})
}