package main

import "fmt"

// Struct definition
type Person struct {
	First string
	Last  string
	Age   int
}

//DoubleZero is mine
type DoubleZero struct {
	//anonymous type. DoubleZero has directly access to Person fields. Like inheritance
	Person
	First         string
	LicenseToKill bool
}

// We attach this method to Person type using the receiver notation
func (p Person) fullName() string {
	return p.First + " " + p.Last
}

/*
	Overriding methods
*/

// Gretting person
func (p Person) Gretting() {
	fmt.Println("I'm just a regular person")
}

// Greeting dz sample
func (dz DoubleZero) Gretting() {
	fmt.Println("Miss Moneypenny, so good to see you")
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   31},
		// The outer field First overrides the inner type.
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}
	p2 := Person{"Miss", "Moneypenny", 25}

	fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenseToKill)
	fmt.Println(p2.First, p2.Last, p2.Age)
	// Person and DZ have the same field First. We can't inherance this field
	fmt.Println(p1.First, p1.Person.First)
	// We can access Age field without use person, only DoubleZero level
	fmt.Println(p1.Age)
	fmt.Println(p1.fullName())

	p1.Gretting()
	p2.Gretting()

	// PONTERS WITH STRUCTS
	p3 := &Person{"M", "commander", 60}

	fmt.Println(p3)
	fmt.Printf("%T\n", p3)
	fmt.Println(p3.First)
	fmt.Println(p3.Age)
}
