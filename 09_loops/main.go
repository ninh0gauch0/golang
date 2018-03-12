package main

import "fmt"

func main() {

	// init-condition-post
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	// nested loops
	fmt.Println("Nested loops")

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}

	// while
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// for break
	i = 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}

	// for continue
	i = 0
	for {
		i++
		if i%2 == 0 {
			// Skips the rest of iteration and goes to next iteration
			continue
		}

		fmt.Println(i)
		if i >= 50 {
			break
		}

	}

	// Runes (char) and strings
	for i := 50; i <= 140; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}

	foo := 'a'
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo))
}
