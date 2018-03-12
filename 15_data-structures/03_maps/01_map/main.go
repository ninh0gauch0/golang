package main

import (
	"fmt"
)

func main() {

	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)
	// m := map[string]string{}
	// Set key/value pairs using typical name[key] = val syntax.
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	//Get a value for a key with name[key]
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// The builtin delete removes key/value pairs from a map.
	delete(m, "k2")
	fmt.Println("map:", m)

	/*
		Sometimes you need to distinguish a missing entry from a zero value.
		Is there an entry for "UTC" or is that 0 because it's not in the map at all?
		You can discriminate with a form of multiple assignment.

		The optional second return value when getting a value from a map indicates
		if the key was present in the map. This can be used to disambiguate between
		missing keys and keys with zero values like 0 or "".

		Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	*/
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	myGreeting := map[int]string{
		0: "Hola caranchoa!",
		1: "Hi de puta!",
		2: "Que te den",
		3: "What a shit!",
	}

	fmt.Println(myGreeting)

	if val, exists := myGreeting[2]; exists {
		delete(myGreeting, 2)
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	// First var is key and the second one is the value
	for k, v := range myGreeting {
		fmt.Println(k, "-", v)
	}

	// Like with arrays there is also a shorter way to create maps
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elements)

}
