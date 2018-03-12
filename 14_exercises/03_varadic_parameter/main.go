package main

import (
	"fmt"
)

func max(numbers ...int) int {
	var max int
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	greatest := max(15, 56, 99, -5, 19, 86, 534534, 76457456, 7624, 42134, -523523)
	fmt.Printf("The greatest number is: %d", greatest)
}
