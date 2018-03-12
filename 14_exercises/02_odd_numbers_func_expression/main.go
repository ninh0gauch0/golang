package main

import "fmt"

func main() {

	var number int

	half := func(number int) (int, bool) {
		return number / 2, number%2 == 0
	}

	fmt.Println("Hi, please, introduce a number")
	fmt.Scanf("%d", &number)

	resp, even := half(number)

	fmt.Printf("The half value is %d \n", resp)
	fmt.Printf("Is it even? %t \n", even)
}
