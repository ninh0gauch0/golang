package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	//This produces a deadlock
	for {
		fmt.Println(<-c)
	}
}
