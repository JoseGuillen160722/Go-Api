package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Registrando la ruta y el handler
	http.HandleFunc("/saludar", saludar)
	http.ListenAndServe(":8080", nil)
}

func saludar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo")
}
