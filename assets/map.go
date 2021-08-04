package main

import "fmt"

// Tanto los json, los diccionarios y los maps no son mas que [Hash Map]
// Los MAP usan concurrencias de forma nativa

func main() {
	m := make(map[string]int)
	m["Tomi"] = 22
	m["Guido"] = 16

	fmt.Println(m)

	// Recorrer Map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor
	value := m["Tomi"]
	fmt.Println(value)

	value2, ok := m["Tito"] // El ok es para que tire true or false si existe o no la key en el map
	fmt.Println(value2, ok)
}
