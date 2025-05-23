package main

import "fmt"

func main() {
	/*
		Es similar a un objeto en js o un diccionario en python
		la unica diferencia es que se espeicifca el tipo de dato
		de la llave, como tambien de los valores.
	*/
	dias := make(map[string]int)

	dias["Domingo"] = 0
	dias["Lunes"] = 1
	dias["Martes"] = 2

	fmt.Println(dias["Domingo"])
}
