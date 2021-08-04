package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

// el primer parámetro es lo que se ingresa, el segundo la cantidad de returns que tendrá. En ambos casos se especifica el tipo de dato.
func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola Mundo")
	tripleArgument(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("Value:", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("value1 y value2:", value1, value2)

	// En una función que retorna dos valores, se puede seleccionar de la siguiente manera cuántos valores queremos retornar en este caso específico al llamarla
	value3, _ := doubleReturn(3)
	fmt.Println("value3:", value3)
}
