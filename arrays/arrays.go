package main

import "fmt"

func main() {
	// Con los valores por defecto
	var numbers [7]int

	fmt.Println(numbers)

	// Sin los valores por defecto al inicio
	names := [2]string{"Mike", "Chloe"}

	fmt.Println(names)
	// Longitud del array útil en los bucles
	// Longitud de array fija al iniciarla
	fmt.Println(len(names))

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Arrays bidimensionales, tridimensionales, ...
	var matriz [7][5]int

	fmt.Println(matriz)

	matriz[1][2] = 7

	fmt.Println(matriz)
	fmt.Println(matriz[1][2])
}
