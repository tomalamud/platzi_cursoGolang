package main

import "fmt"

func main() {
	// Defer: antes de morir
	defer fmt.Println("Hola, esto siempre estará al final")
	fmt.Println("Esto estará normalmente al principio")

	// Continue y break necesitan un for
	for i := 0; i < 10; i++ {
		if i != 2 {
			fmt.Println(i)
		}

		// Continue:
		if i == 2 {
			fmt.Println("Es un 2!!")
		}

		// Break:
		if i == 8 {
			break
		}
	}
}
