package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := 24
	edad_str := strconv.Itoa(edad)
	// No puede concatenar un nº con una cadena.
	fmt.Println("Mi edad es: " + edad_str)

	age := "24"
	// age_str,err := strconv.Atoi(age)
	// Da error al no usar variable err.
	// Operador _ (guión bajo) es como una variable que se usa cuando no se va a usar.
	age_str, _ := strconv.Atoi(age)
	fmt.Println(age_str + 12)
}
