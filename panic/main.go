package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Ingrese el dividendo: ")
	fmt.Scanf("%d", a)
	fmt.Print("Ingrese el divisor: ")
	fmt.Scanf("%d", b)

	if b == 0 {
		panic("No es posible didivir por Zero")
	}

	resultado := a / b
	fmt.Print(resultado)

}
