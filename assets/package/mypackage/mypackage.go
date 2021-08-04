package mypackage

// CarPublic "Crea un Car público"
type CarPublic struct {
	Brand string
	Year  int
	wheel string // este método es privado, comienza con minúscula
}

// Que los métodos, las funciones o los structs sean públicos o privados quiere decir que se exportarán o no desde el paquete.
type carPrivate struct {
	brand int
}
