package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p\n", &z) // address in zero func
	fmt.Println(&z)
	z = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x)
	zero(x)
	fmt.Println(x) // x is still 5
}
