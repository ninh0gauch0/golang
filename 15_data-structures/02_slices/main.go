package main

import (
	"fmt"
)

func main() {
	// WITH VAR. No number inside the brackets. Has NIL VALUE.
	var oneSlice []int
	//SHORTHAND. Explicity included the data
	mySlice := []int{1, 3, 5, 7, 9, 11}
	// WITH MAKE
	mySlice2 := make([]bool, 2, 4)

	fmt.Println("%T - %T - %T\n", mySlice, mySlice2, oneSlice)
	fmt.Println(mySlice)

	/*
		Unlike arrays, slices are typed only by the elements they contain
				(not the number of elements). To create an empty slice with non-zero length,
				use the builtin make.

				Here we make a slice of strings of length 3 (initially zero-valued).

				The second parameter is the slice's lenght and the third one is the capacity of the slice.

				The capacity is the number of elements for which there is space allocated in the underlying array.
				At any time the following relationship holds:

				0 <= len(s) <= cap(s)

				If we pass the capacity of a slice, a new memory allocation was created with 2xcap, deleting the old one.

				For a given nil slice or channel, the length and the capacity are 0.
	*/
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	/*
		Slices support several more that make them richer than arrays.
		One is the built in append, which returns a slice containing
		one or more new values.

		Note that we need to accept a return value from append as we may get
		a new slice value.
	*/
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	/*
		Slices can also be copy’d. Here we create an empty slice c of the same
		length as s and copy into c from s.
	*/
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	/*
		Slices support a “slice” operator with the syntax slice[low:high].
		For example, this gets a slice of the elements s[2], s[3], and s[4].
	*/
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to (but excluding) s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("sl3:", l)

	// We can declare and initialize a variable for slice in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	/*
		Playing with slice capacity

		If we dont inform the capacity. First is one, then 2, then 4, then 8, ...
	*/

	coolSlice := make([]int, 0, 5)
	fmt.Println("-------------------------")
	fmt.Println(coolSlice)
	fmt.Println(len(coolSlice))
	fmt.Println(cap(coolSlice))
	fmt.Println("-------------------------")

	for i := 0; i < 80; i++ {
		coolSlice = append(coolSlice, i)
		fmt.Println("Len: ", len(coolSlice), "Cap: ", cap(coolSlice), "Value: ", coolSlice[i])
	}

	/*
		Slices can be composed into multi-dimensional data structures.
		The length of the inner slices can vary, unlike with multi-dimensional arrays.
	*/
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// PLaying Multi dimensional slices

	records := make([][]string, 0)

	//student 1
	student1 := make([]string, 4, 8)
	student1[0] = "Foster"
	student1[1] = "Natham"
	student1[2] = "100.00"
	student1[3] = "74.00"

	//store the data
	records = append(records, student1)

	//student 2
	student2 := make([]string, 4, 8)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.00"
	student2[3] = "96.00"

	records = append(records, student2)

	fmt.Println(records)
}
