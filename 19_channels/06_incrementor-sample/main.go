package main

import (
	"fmt"
)

func main() {
	s1 := "Foo:"
	s2 := "Bar:"

	c1 := incrementor(s1)
	c2 := incrementor(s2)
	c3 := puller(c1, s1)
	c4 := puller(c2, s2)
	fmt.Println("Final counter:", <-c3+<-c4)
}

func incrementor(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- 1
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c chan int, s string) chan int {
	out := make(chan int)

	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		fmt.Println("Partial count for "+s, sum)
		out <- sum
		close(out)
	}()
	return out
}
