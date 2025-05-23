package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//Iniciar la semilla basada en la hora actual
	rand.Seed(time.Now().UnixNano())

	var calificacion = rand.Intn(10)
	switch {
	case calificacion == 10:
		fmt.Println("Excelente")
	case calificacion < 10 && calificacion >= 1:
		fmt.Println("Deberias esforzarte mas!")
	default:
		fmt.Println("valor invalido ingresado")
	}

	switch calificacion {
	case 10, 8, 7, 6:
		fmt.Println("Excelente")
	case 5, 4, 3, 2, 1:
		fmt.Println("Deberias esforzarte mas!")
	default:
		fmt.Println("valor invalido ingresado")
	}

}
