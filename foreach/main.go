package main

import "fmt"

func main() {
	animales := [...]string{"perro", "gato", "etc", "etc3"}

	for indice := 0; indice < len(animales); indice++ {
		elemento := animales[indice]
		fmt.Println(elemento)
	}

	for indice, elemento := range animales {
		fmt.Println("El indice es: ", indice, " el valor es: ", elemento)
	}
}
