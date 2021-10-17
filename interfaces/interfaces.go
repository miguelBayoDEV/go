package main

import "fmt"

// Interface es estructura con métodos vacíos o no implementados
type User interface {
	Permissions() int // Ej: 1-5
	Name() string
}

// Struct que se le implementa la interface
type Admin struct {
	name string
}

// Métodos de la interface para struct Admin
func (this Admin) Permissions() int {
	return 5 // Tiene todos los permisos
}

func (this Admin) Name() string {
	return this.name
}

// Otro struct que se le implementa la interface
type Editor struct {
	name string
}

// Métodos de la interface para struct Editor
func (this Editor) Permissions() int {
	return 3 // Tiene todos los permisos
}

func (this Editor) Name() string {
	return this.name
}

// Función genérica para autentificar autorización
func auth(user User) string {
	if user.Permissions() > 4 {
		return user.Name() + " tiene permisos de administrador"
	} else {
		return user.Name() + " no tiene permisos de administrador"
	}
}

func main() {
	admin := Admin{"Mike"}

	fmt.Println(auth(admin))

	editor := Editor{"Josh"}

	fmt.Println(auth(editor))

	// Otra forma
	users := []User{Admin{"Elena"}, Editor{"Raquel"}}

	for _, user := range users {
		fmt.Println(auth(user))
	}
}
