package main

import "fmt"

func main() {
	age := 44
	fmt.Println(&age) // Prints an address

	changeMe(&age)

	fmt.Println(&age) // Same address
	fmt.Println(age)  // 24 with the pointer function, 44 with the pass by value function
}

func changeMe(z *int) {
	fmt.Println(z)  // Same address
	fmt.Println(*z) // 44
	*z = 24
	fmt.Println(z)  // Same address
	fmt.Println(*z) // 24
}

// when changeMe is called in line 9 the value
// 44 is being passed as an argument
func changeMePassByValue(z int) {
	fmt.Println(z) // 44
	z = 24
	fmt.Println(z) // 24
}
