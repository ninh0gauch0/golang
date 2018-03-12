package main

type Animal struct {
	sound string
}

type Dog struct {
	Animal
	friendly bool
}

type Cat struct {
	Animal
	annoying bool
}
