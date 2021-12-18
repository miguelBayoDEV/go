package main

import "fmt"

func main() {
	// Channels permiten comunicarse entre varias gorrutines

	// Primero definir un canal
	// Primer dato chan para saber que el dato creado es un canal y el segundo tipo dato o estructuras que se enviará entre canales
	channel := make(chan string)

	// Crear una gorrutine y le paso el canal creado
	// Gorrutine en ciclo infinito
	go func(channel chan string) {
		// Escanear datos
		for {
			var name string
			fmt.Scanln(&name)

			// Forma mandar información de un canal es: nombreCanalAEnviar '<-' más dato
			channel <- name
		}
	}(channel)

	// Nueva variable que recibe información del canal es al revés: dato <- canal
	/*
		Otra forma:
		var infoRecibida string
		infoRecibida <- channel
	*/
	infoRecibida := <-channel

	// Imprimer lo que se recibe del canal
	fmt.Println("Estoy recibiendo info canal primera: " + infoRecibida)

	infoRecibida = <-channel

	fmt.Println("Estoy recibiendo info canal segunda: " + infoRecibida)

	// Bucle infinito para imprimir información
	/*
		infoRecibida := <-channel
		fmt.Println("Estoy recibiendo info canal primera: " + infoRecibida)
	*/
}
