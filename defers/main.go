package main

import (
	"fmt"
	"os"
)

func main() {
	//La palabra reservada defer,
	// permite que se ejecute una función al final
	// de la ejecución de la función que la contiene.
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Se ha producido un error:", err)
		}
		fmt.Println("Aqui iria la logica despues del error")
	}()

	file, err := os.Open("examples.txt")

	if err != nil {
		panic(err)
	}

	defer func() {
		fmt.Println("Closing file...")
		file.Close()
	}()

	contenido := make([]byte, 254)
	longitud, _ := file.Read(contenido)

	contenidoArchivo := string(contenido[0:longitud])
	fmt.Println("Contenido del archivo:", contenidoArchivo)
}
