package main

import "fmt"

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println(areaCuadrado)

	// OPERADORES
	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = y - x
	fmt.Println("Resta:", result)

	// Multiplicacion
	result = y * x
	fmt.Println("multiplicacion:", result)

	// Division
	result = y / x
	fmt.Println("division:", result)

	// Modulo
	result = y % x
	fmt.Println("modulo:", result)

	// Incrementar
	x++
	fmt.Println("Incremental", x)

	// Decremental
	x--
	fmt.Println("Decremental", x)

	// RETOS
	// Rectangulo, trapecio, circulo
}
