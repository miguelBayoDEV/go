package main

import "fmt"

func main() {
	x := 7
	y := 3

	/**
		== igual a
		!= diferente de
		>  mayor que
		<  menor que
		>= mayor o igual que
		<= menor o igual que
		&& AND
		|| OR
	**/

	if x > y {
		fmt.Printf("%d es mayor que %d\n", x, y)
	} else if x < y {
		fmt.Printf("%d es menor que %d\n", x, y)
	} else {
		fmt.Println("Son iguales")
	}
}
