package main

import (
	"fmt"
)

func makeGreeter() func() string {
	return func() string {
		return "Hello madafakah!"
	}
}

func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	greeting := func() {
		fmt.Println("Pollas in vinegar")
	}

	greeting()
	fmt.Printf("%T\n", greeting)

	// A func expression that returns a function
	greet := makeGreeter()
	fmt.Println(greet())

	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())

}
