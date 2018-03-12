package main

import (
	"fmt"
)

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // Prints memory dir

	// This is a pointer to an int variable
	var b *int = &a // reference

	fmt.Println(b)  // Prints memory dir
	fmt.Println(*b) // 43. Dereference

	*b = 69        // Cmabiamos el valor de esa dirección de memoria, cambiándolo a 69
	fmt.Println(a) // 69

}
