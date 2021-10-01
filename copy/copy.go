package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	copia := make([]int, len(slice))

	fmt.Println(slice)
	fmt.Println(copia)

	// Función copy permite copiar dos array
	// Copy copia la mínima longitud de los arrays
	// Estructura: copy(destino array, fuente array)

	copy(copia, slice)

	fmt.Println(slice)
	fmt.Println(copia)
}
