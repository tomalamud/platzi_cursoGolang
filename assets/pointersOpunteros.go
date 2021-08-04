package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() { // antes de la func puse el puntero
	fmt.Println(myPc.brand, "Pong")
}

func (myPc *pc) duplicateRAM() {
	myPc.ram = myPc.ram * 2
}

func main() {

	a := 50
	b := &a

	fmt.Println(a)
	fmt.Println(b) // dirección de memoria -> '0xc000014088'
	fmt.Println(*b)

	*b = 100
	fmt.Println(b)
	fmt.Println(a)

	myPc := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPc)

	myPc.ping()

	myPc.duplicateRAM()
	fmt.Println(myPc)
	myPc.duplicateRAM()
	fmt.Println(myPc)
}
