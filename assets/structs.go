package main

import "fmt"

type car struct {
	brand string
	year  int
	wheel string
}

func main() {
	// forma #1 de instanciar clase
	newCar := car{brand: "Mercedes Benz", year: 2021, wheel: "Good Year"}
	fmt.Println(newCar)

	// forma #2 de instanciar clase
	var otherCar car
	otherCar.brand = "ferrari"
	fmt.Println(otherCar)
}
