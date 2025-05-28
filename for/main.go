package main

import "fmt"

func main() {
	usuarios := map[int]string{}

	usuarios[1] = "Gabykap"
	usuarios[2] = "Marykap"
	usuarios[3] = "Emmitakap"

	for llave, valor := range usuarios {
		fmt.Println(llave, valor)
	}
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
