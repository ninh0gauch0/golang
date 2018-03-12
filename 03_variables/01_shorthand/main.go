package main

import "fmt"

var g = "cowboy"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true

	// Print values

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	// Print types
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)

}
