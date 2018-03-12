package main

import "fmt"

func main() {
	x := 1
	str := evalInt(x)
	fmt.Println(str)
}

func evalInt(n int) string {
	if n > 10 {
		return fmt.Sprint("x is greater than 10")
	} else { // We must delete this else statement to pass lint review
		return fmt.Sprint("x is lesser than 10")
	}
}
