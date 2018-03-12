package main

import (
	"fmt"
)

// Empty interface
type vehicles interface{}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

// This function receives an empty interface
func specs(a interface{}) {
	fmt.Println(a)
}

func main() {

	prius := car{}
	tacoma := car{}
	bmw528 := car{}
	boeing747 := plane{}
	boeing757 := plane{}
	boeing767 := plane{}
	sanger := boat{}
	malibu := boat{}

	rides := []vehicles{prius, tacoma, bmw528, boeing747, boeing757, boeing767, sanger, malibu}

	for k, v := range rides {
		fmt.Println(k, " - ", v)
	}

	// We have to exec with `go run *.go`
	fido := Dog{Animal{"woof"}, true}
	fifi := Cat{Animal{"meow"}, true}
	specs(fido)
	specs(fifi)

}
