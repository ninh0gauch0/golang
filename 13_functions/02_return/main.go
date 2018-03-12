package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe "))
	fmt.Println(greet2("Fuck ", "You "))
	a, b := greet3("Die", "madafakah")
	fmt.Println(a, b)
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

/*
IMPORTANT
Avoid using named returns.

Occasionally named returns are useful. Read this article for more information:
https://www.goinggo.net/2013/10/functions-and-naked-returns-in-go.html
*/

func greet2(fname string, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}

/* Return many*/

func greet3(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
