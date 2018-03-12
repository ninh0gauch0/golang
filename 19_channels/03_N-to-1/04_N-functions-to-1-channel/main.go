package main

import (
	"fmt"
)

func main() {
	//Number of times we will loop
	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	// On this stuff we are picking channel values
	for n := range c {
		fmt.Println(n)
	}
}
