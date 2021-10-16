package main

// Go no está orientado a objeto y esto hace una estructura de lenguaje orientado a objetos y herencia

import "fmt"

// Herencia de structs
type Employer struct {
	name string
}

type Person struct {
	name string
}

// Si hereda de dos struct se deberá poner object.nombreStruct.propiedad sino da error de ambigüedad
type Programmer struct {
	// Campo anónimo
	Employer
	Person
}

// Método de Employer herencia padre
func (this Employer) speak() string {
	return "Good Morning!!!"
}

// Sobreescribir método heredado de Employer y no hay error de ambigüedad.
func (this Programmer) speak() string {
	// Concatenar el método padre con el sobreescrito
	return this.Employer.speak() + " - Good Morning, boss!!!"
}

func main() {
	// Instanciar programmer
	programmer := Programmer{Employer{"Mike"}, Person{"Humm"}}

	// Error ambigüedad
	// fmt.Println(programmer.name)

	// Otra forma
	fmt.Println(programmer.Employer.name)
	fmt.Println(programmer.Person.name)

	// Imprimir método
	fmt.Println(programmer.speak())

}
