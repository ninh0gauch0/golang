package server

import (
	"context"
	"fmt"
	"time"
)

//IncrementCounter is a shit
func (w *Worker) incrementCounter() {
	w.logger.Infof("Incrementing a counter")
	w.counter = w.counter + 1
	fmt.Printf("Counter's value is %d \n", w.counter)
}

func (w *Worker) start(ctx *context.Context) bool {
	w.ticker = time.NewTicker(w.Period)
	w.counter = 0
	w.logger.Infof("Worker started")

	go func(ticker *time.Ticker) {
		for {
			// Preguntar al teacher
			select {
			case <-(*ctx).Done():
				w.logger.Infof("Worker stopped")
				ticker.Stop()
				w.logger.Infof("I was able to tick: %d times", w.counter)
				// Teacher
			case <-ticker.C:
				w.incrementCounter()
			}
		}
	}(w.ticker)

	return true
}
