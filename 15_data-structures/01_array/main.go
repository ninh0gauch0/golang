package main

import "fmt"

func main() {
	var x [78]string
	// It prints the array with default string value " "
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 65; i <= 122; i++ {
		// string(number) = ascii code
		x[i-65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	var x2 [256]int

	fmt.Println(len(x2))
	// Prints int default value 0
	fmt.Println(x2[42])

	for i := 0; i < 256; i++ {
		x2[i] = i
	}

	for i, v := range x2 {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}

	var x3 [256]int
	fmt.Println(len(x3))
	// Prints byte default value 0
	fmt.Println(x3[42])

	for i := 0; i < 256; i++ {
		x3[i] = i
	}

	for i, v := range x3 {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}
}
