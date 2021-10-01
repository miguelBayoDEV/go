package main

import "fmt"

func main() {
	// Parámetros función make son: slice, longitud y capacidad del slice
	slice := make([]int, 3, 7)
	fmt.Println(slice)

	// Append agrega al slice un nuevo elemento
	// Al superar la capacidad del slice con método append reconstruye el slice
	slice = append(slice, 7)

	fmt.Println(slice)
	fmt.Println(len(slice))

	// Función cap de capacidad del slice
	fmt.Println(cap(slice))
}
