package main

import "fmt"

func main() {
	/*
		&&
		||
		!
	*/

	and := 20 == 20 && !(30 > 20)
	fmt.Println(and)
	or := 20 == 20 || !(30 > 20)
	fmt.Println(or)
	not := !false
	fmt.Println(not)
}
