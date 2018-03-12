package main

import "fmt"

// Package level scope
var x int = 42

func main() {
	fmt.Println(x)
	foo()

	// Declaration and Assignation. LOCAL SCOPE
	y := 17
	fmt.Println(y)
}

func foo() {
	fmt.Println(x)
}
