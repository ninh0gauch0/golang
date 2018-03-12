package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("your_mamma.txt")

	if err != nil {
		fmt.Println("err happened", err)

		/*
			Package log implements a simple logging package... writes to
			standard error and prints the date and tim of each logged message.

			The Fatal functions call os.Exit(1) after writing the log message.

			Panic functions call panic after writing the log message.
		*/
		log.Println("err happened 2.0", err)
		log.Fatalln(err)
		panic(err)
	}
}
