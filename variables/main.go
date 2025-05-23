package main

import (
	"fmt"
	"reflect"
)

// Declarar variables especificando el tipo de dato
var saludo string = "Hola Mundo"

func main() {

	//declarar variables sin especificar el tipo, ni var
	//Solo es posible usar dentro de una funcion.
	saludo2 := "Saludos!"
	// tambien es posible sin especificar el tipo y usando var
	var saludo3 = "Saludando*"

	fmt.Println("saludo 1: ", saludo, "saludo 2: ", saludo2, "saludo 3: ", saludo3)

	/*Declarar varias variables en una linea
	 */

	var nombre, apellido, pais string = "Gabriel", "Acosta", "Argentina"

	nombreDos, apellidoDos, edad := "Mariela", "Gonzalez", 25
	//apellido[2] esto devuelve el valor en ASCI
	fmt.Println(len(nombre), apellido[2], pais)
	fmt.Println(nombreDos, apellidoDos, edad)

	fmt.Println(reflect.TypeOf(apellido[2]))

	//Para podes imprimir solo texto debo apoyarme desde PrintF
	fmt.Printf("%c\n ", apellido[2])

}
