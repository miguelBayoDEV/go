package main

import "fmt"

// Type es una palabra clave que hace referencia a un nuevo tipo de dato
type User struct {
	// Propiedades o atributos struct
	age            int
	name, lastName string
}

// Métodos. Su sintasis:
func (usuario User) nombreCompleto() string {
	return usuario.name + " " + usuario.lastName
}

// Para hacer un setter y que modifique el valor de una propiedad de struct se usa el puntero sino haría una copia
func (this User) set_name(n string) {
	this.name = n
}

func (this *User) setName(n string) {
	this.name = n
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
	mig.lastName = "Bay"
	(*mig).age = 25

	fmt.Println(mig)
	fmt.Println(mig.name)
	fmt.Println((*mig).age)

	// Imprimir método
	fmt.Println(mig.nombreCompleto())

	// No modifica nada porque hace una copia
	mig.set_name("Door")
	fmt.Println(mig.name)

	// Al ser un puntero se imprime la modificación
	mig.setName("Jeremo")
	fmt.Println(mig.name)
}
