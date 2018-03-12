package main

import (
	"fmt"
)

func main() {

	var x = 12
	var y = 12.1230123
	// conversion int to float
	fmt.Println(y + float64(x))
	// float to int
	fmt.Println(int(y) + x)

	var a rune = 'a' // rune is an alias for int32
	var b int32 = 'b'
	fmt.Println(a)
	fmt.Println(b)
	// rune to string
	fmt.Println(string(a))
	//int32 to string
	fmt.Println(string(b))

	var name interface{} = "Sydney"
	str, ok := name.(string)

	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Println("Value is not a string")
	}

	var val interface{} = 7
	// fmt.Println(val + 6) // This doesn't work
	fmt.Println(val.(int) + 6)

	rem := 7.24
	fmt.Printf("%T\n", rem)
	fmt.Printf("%T\n", int(rem)) //THIS IS CASTING!

	fmt.Printf("%T\n", val)
	//fmt.Printf("%T\n", int(val)) // This fails. We can not convert type interface with int. We need assertion.
	fmt.Printf("%T\n", val.(int))
}
