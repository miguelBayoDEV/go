package main

import "fmt"

//import "fmt"

func main() {
	// Siempre que se generen variables se deberán declarar sino no compilará. Tipado dinámico go.
	// Fuertemente típado no se puede pasar un nº a una cadena y viceversa.
	x := 7
	fmt.Println(x)
	nombre := "Miguel"
	fmt.Println(nombre)
	// Cambiar variable valor de variable ya creada sino habrá que crear una nueva.
	nombre = "José"
	fmt.Println(nombre)
}
