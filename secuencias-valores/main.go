package main

import "fmt"

const (
	//iota siempre comienza desde 0, en este caso,
	// sirve para crear una secuencia a partir de 0 dentro de las
	//declaracion de contantes
	Domingo int = iota + 5
	Lunes
	martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println(Domingo)
	fmt.Println(Lunes)
	fmt.Println(martes)
	fmt.Println(Miercoles)
	fmt.Println(Jueves)
	fmt.Println(Viernes)
	fmt.Println(Sabado)
}
