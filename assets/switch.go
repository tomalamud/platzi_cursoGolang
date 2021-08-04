package main

import "fmt"

func main() {
	modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Sin condicion
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es menor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Print("No condicion")
	}
}
