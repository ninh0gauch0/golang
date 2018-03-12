package main

import "fmt"

// Some types are referenced types and dont need to use pointers.

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Todd"])
}

func changeMe(z map[string]int) {
	z["Todd"] = 44
	fmt.Println(z)
}
