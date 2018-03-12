package main

import "fmt"

const (
	_        = iota             // 0. Represents successive untyped integer constants
	KB       = 1 << (iota * 10) // 1 << (1 * 10) Desplazamos 10 posiciones los bits (1024 bytes)
	MB       = 1 << (iota * 10) // 1 << (2 * 10) Desplazamos 20 posiciones los bits (1024 Kb)
	GB       = 1 << (iota * 10) // 1 << (3 * 10) Desplazamos 30 posiciones los bits (1024 Mb)
	TB       = 1 << (iota * 10) // 1 << (4 * 10)
	Pi       = 3.14
	Language = "Go"
)

func main() {
	fmt.Println(Pi)
	fmt.Println(Language)
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}
