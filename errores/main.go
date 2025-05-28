package main

import (
	"errors"
	"fmt"
)

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No es posible dividir por 0")

	} else {
		return dividendo / divisor, nil
	}
}

// Nil hace referencia a ningun valor como error
func main() {
	if resultado, err := division(10, 0); err != nil {
		panic(err)
	} else {
		fmt.Println("el resultado es: ", resultado)
	}
}
