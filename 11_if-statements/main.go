package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("This ran")
	}

	if false {
		fmt.Println("This did not run")
	}

	// Initialization statement
	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	/* 	if err := file.Chmod(0664); err != nil {
		log.Print(err)
		return err
	} */

	// if - else if - else
	if false {
		fmt.Println("first print statement")
	} else if false {
		fmt.Println("second print statement")
	} else if true {
		fmt.Println("ahahaha print statement")
	} else {
		fmt.Println("third print statement")
	}
}
