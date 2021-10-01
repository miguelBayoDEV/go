package main

import "fmt"

func main() {
	// No tiene valor fijo al iniciarla = slices
	// Primera forma de declarar un slice
	var slices []int

	if slices == nil {
		fmt.Println("It's empty!!!")
	} else {
		fmt.Println(slices)
	}

	//Segunda forma de declarar un slice
	slice := []int{7, 5, 3, 1}
	fmt.Println(slice)
	// len() para longitud del slice
	fmt.Println(len(slice))

	// Partir array en slice (slicing)
	array := [7]int{7, 6, 5, 4, 3, 2, 1}
	sl := array[2:5]
	fmt.Println(sl)
}
