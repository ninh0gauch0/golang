package main

import (
	"fmt"
)

func main() {
	// Channel creation
	c := make(chan int)
	//Channel used as semaphore
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	// This stuff has to be included into a go rutine.
	// If we add this code only into the main func, we'll have a race condition
	go func() {
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
