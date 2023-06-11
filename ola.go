package main

import "fmt"

const prefixHelloPortuguese = "Olá, "

func Ola(name string) string {
	if name == "" {
		name = "Mundo"
	}
	
	return prefixHelloPortuguese + name
}

func main() {
	fmt.Println(Ola("Eric"))	
}
