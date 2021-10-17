package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//mySlowName("Michael")

	// Uso de go delante del método para que se ejecute en otro plano y sea concurrente
	go mySlowName("Mike")

	// Código anterior es sincrono y se debe esperar a que acabe el anterior
	fmt.Println("Boring!!!")

	// Para que no acabe el programa hasta que pulse una tecla
	var wait string
	fmt.Scanln(&wait)
}

// Función que parte la letras con split de strings
func mySlowName(name string) {
	letras := strings.Split(name, "")

	for _, letra := range letras {
		time.Sleep(2000 * time.Millisecond)
		fmt.Println(letra)
	}
}
