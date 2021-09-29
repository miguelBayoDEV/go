package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	method := "método"
	fmt.Print("Buenas!!")
	// Hacer salto de línea
	// fmt.Println()
	fmt.Println("El más usado " + method)
	// Otra forma de imprimir datos pasándole variables también. %d número, %v ó %s cadena, %t boolean ...
	fmt.Printf("El más usado %v", method)
	// fmt.Println()
	age := 24
	fmt.Printf("Mi edad es: %d", age)
	// fmt.Println()

	//Leer datos
	var age_read int
	fmt.Println("Introduzca tu edad: ")
	// Scanf mismos valores que Prinf con \n salto de línea y poner & delante de la variable
	fmt.Scanf("%d\n", &age_read)
	fmt.Printf("Mi edad después de leer la variable es : %d\n", age_read)

	var name string
	fmt.Println("Introduzca el nombre: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Bienvenido %s\n", name)

	// Otra forma es crear un reader importando bufio y os
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Introduzca tu nombre completo: ")
	// Devuelve el nombre y la comprubación de si hay un error o no. ReadString leer hasta el carácter \n en comillas simples
	nombre, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bienvenido " + nombre)
	}
}
