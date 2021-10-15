package main

import "fmt"

// Type es una palabra clave que hace referencia a un nuevo tipo de dato
type User struct {
	// Propiedades o atributos struct
	age            int
	name, lastName string
}

func main() {
	var mike User

	fmt.Println(mike)

	// Otra forma de declarar la struct
	fmt.Println(User{age: 24, name: "Mike", lastName: "Bayo"})

	// Otra forma
	josh := User{age: 29, name: "Josh", lastName: "Door"}

	fmt.Println(josh)

	// Otra forma con keywork new lo asigna como puntero
	mig := new(User)

	// Structs son mutables por lo que se puede variar los valores de los campos después de instanciarlos
	mig.name = "Mig"
	(*mig).age = 25

	fmt.Println(mig)
	fmt.Println(mig.name)
	fmt.Println((*mig).age)
}
