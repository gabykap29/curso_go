package main

import "fmt"

func main() {
	arregloUno := [...]int{1, 2, 3, 4, 5, 6, 7}
	var arregloDos [10]int
	monedas := [...]string{0: "Dolar", 1: "Peso", 3: "Euro"}

	for i := 0; i < 10; i++ {
		arregloDos[i] = i + 10
	}
	fmt.Println(arregloUno)
	fmt.Println(arregloDos)
	fmt.Println(monedas)
}
