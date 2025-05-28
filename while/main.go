package main

import "fmt"

func main() {
	//Go no tiene while, pero es posible emularlo.

	numero := 12345
	contador := 0

	for numero > 0 {
		numero = numero / 10
		contador++
		fmt.Println("Cantidad de digitos ", numero)
		fmt.Println(contador)
	}

}
