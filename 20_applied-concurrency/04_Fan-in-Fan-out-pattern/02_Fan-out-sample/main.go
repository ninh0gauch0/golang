package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

	// FAN OUT
	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// FAN IN
	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	wg.Add(len(cs))

	// Start an output goroutine for each input channel in cs.
	// output copies values from c to out until c is closed, then calls wg.Done.

	/*
		output := func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}

		for _, c := range cs {
			go output(c)
		}

	*/

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

/*
FAN OUT
Multiple functions reading from the same channel until that channel is closed

FAN IN
A function can read from multiple inputs and proceed until all are closed by
multiplexing the input channels onto a single channel that's closed when
all the inputs are closed.

PATTERN
there's a pattern to our pipeline functions:
-- stages close their outbound channels when all the send operations are done.
-- stages keep receiving values from inbound channels until those channels are closed.

source:
https://blog.golang.org/pipelines
*/
