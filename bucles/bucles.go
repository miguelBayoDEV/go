package main

import "fmt"

func main() {

	// Bucle for
	for i := 0; i < 7; i++ {
		fmt.Println(i)
	}

	// Bucle while
	a := 0
	for {
		if a == 4 {
			a++
			continue
		}

		fmt.Println(a)
		a++

		if a > 10 {
			// Expresión break acaba el bucle y continue cuando se llega a él se vuelve a ejecutar el bucle de nuevo
			break
		}
	}
}
