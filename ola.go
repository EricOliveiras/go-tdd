package main

import "fmt"

const spanish = "spanish"
const french = "french"
const prefixHelloPortuguese = "Olá, "
const prefixHelloEspanish = "Hola, "
const prefixHelloFrench = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "Mundo"
	}

	prefix := prefixOfGreeting(language)

	return prefix + name
}

func prefixOfGreeting(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = prefixHelloEspanish
	case french:
		prefix = prefixHelloFrench
	default:
		prefix = prefixHelloPortuguese
	}

	return
}

func main() {
	fmt.Println(Hello("Eric", ""))
}
