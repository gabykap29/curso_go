package main

import "fmt"

func main() {
	var nombre string
	var edad int
	var altura float32

	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &nombre)

	fmt.Print("Ingresa tu edad: ")
	fmt.Scanf("%d", &edad)

	fmt.Print("Ingresa tu altura: ")
	fmt.Scanf("%f", &altura)

	fmt.Printf("Hola %s, la edad ingresada es %d y tu altura es de %.2f", nombre, edad, altura)

}
