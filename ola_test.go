package main

import "testing"

func TestOla(t *testing.T) {
	result := Ola("Eric")
	expected := "Olá, Eric"

	if result != expected {
		t.Errorf("result: '%s', expected: '%s'", result, expected)
	}
}