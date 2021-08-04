package main

import "fmt"

func main() {
	// Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array,
		len(array), // largo del array
		cap(array)) // capacidad máxima del array

	// Slice: (doc) https://blog.golang.org/slices-intro
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el Slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...) // Descomprime los elementos de la primera lista y los agrega a la segunda
	fmt.Println(slice)

	// Range
	sliceText := []string{"hola", "que", "hace"}
	for i, valor := range sliceText { // si quisieramos mostrar solo el valor sería 'for _, valor :=' | si solo se necesita el indice se pone 'for i :='
		fmt.Println(i, valor)
	}
}
