package main

import (
	"errors"
	"log"
)

func main() {
	_, err := Sqrt(-10)

	if err != nil {
		log.Fatalln(err)
	}

}

//Sqrt function
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		// Creating a custom error
		return 0, errors.New("norgate math: square root of negative number")
	}
	return 42, nil
}
