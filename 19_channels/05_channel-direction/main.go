package main

import (
	"fmt"
)

func main() {
	c := incrementor()
	cSum := puller(c)

	for n := range cSum {
		fmt.Println(n)
	}
}

// Creates the channel, the go routine is executed on background and the channel is returned.
// This channel returned only can be used to receive data (ints)
func incrementor() <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

// This channel passed as parameter only can be used to receive data (ints)
func puller(c <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
