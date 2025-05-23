package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}

	fmt.Println(numeros)

	numeros = append(numeros, 6, 7, 8, 9)
	fmt.Println(numeros)

	nuevoSlice := numeros[0:3]
	fmt.Println(nuevoSlice)
}
