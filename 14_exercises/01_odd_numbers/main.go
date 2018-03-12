package main

import "fmt"

func half(number int) (int, bool) {
	/*	even := false
		res := number / 2

		if number%2 == 0 {
			even = true
		} */

	return number / 2, number%2 == 0
}

func main() {
	var number int

	fmt.Println("Hi, please, introduce a number")
	fmt.Scanf("%d", &number)

	half, even := half(number)

	fmt.Printf("The half value is %d \n", half)
	fmt.Printf("Is it even? %t \n", even)
}
