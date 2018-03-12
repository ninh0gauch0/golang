package main

import "fmt"

func main() {

	fmt.Println(42)

	// Formateo decimal - binary - octal/hexadecimal
	fmt.Printf("%d - %b \n", 42, 42)
	fmt.Printf("%d - %b - %o \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#o \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)

	for i := 0; i < 200; i++ {
		fmt.Printf("%d - %b - %#x \n", i, i, i)
	}
}
