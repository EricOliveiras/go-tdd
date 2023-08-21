package main

import (
	"fmt"
	"io"
	"net/http"
)

func Cumprimenta(escritor io.Writer, nome string) {
	fmt.Fprintf(escritor, "Olá, %s", nome)
}

func HandlerMeuCumprimento(w http.ResponseWriter, r *http.Request) {
	Cumprimenta(w, "Eric")
}

func main() {
	if err := http.ListenAndServe(":8080", http.HandlerFunc(HandlerMeuCumprimento)); err != nil {
		fmt.Println(err)
	}
}
