package main

import "fmt"

func main() {
	// Declaracion variables
	helloMessage := "Hello"
	worldMessage := "world"

	// Println
	fmt.Println(helloMessage)
	fmt.Println(worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos) // asi se agrega el tipo de dato en una impresion
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos) // tipo de dato x default

	// Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Println(message)

	// Tipo de dato de la variable
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
