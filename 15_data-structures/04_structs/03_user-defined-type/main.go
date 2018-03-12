package main

import (
	"fmt"
)

type foo int

func main() {

	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)

	var yourAge int
	yourAge = 18
	fmt.Printf("%T %v \n", yourAge, yourAge)

	// We have to cast our type to can operate with int
	fmt.Println(int(myAge) + yourAge)
}
