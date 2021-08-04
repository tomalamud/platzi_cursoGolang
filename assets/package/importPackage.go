package main

import (
	pk "curso_golang_platzi/assets/package/mypackage" // así se pone alias
	"fmt"
)

// Nótese que el PATH comienza desde el src/ tal como lo configuramos en la variable de entorno $GOPATH (GOMODULES: https://golang.org/doc/tutorial/create-module)

func main() {
	var newCar pk.CarPublic
	newCar.Brand = "Ford"
	newCar.Year = 2020
	fmt.Println(newCar)
}
