package main

import (
	"fmt"
	"strings"
)

func esPalindromo(text string) {
	text = strings.ToLower(text)
	var textReverse string // guardar la palabra en sentido inverso
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i]) // si no ponemos string, el índice nos pasa el ASCII de el caracter
	}

	if text == textReverse {
		fmt.Println("Es palíndromo")
	} else {
		fmt.Println("No es palíndromo")
	}
}

func main() {
	esPalindromo("amoR a roma")
}
