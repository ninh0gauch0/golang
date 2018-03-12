package main

import (
	"fmt"
	"sort"
)

/*This TYPE (person inherits from slice of string) implements sort interface*/
type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func main() {
	// People is a type. We can implement for it the needed methods.
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	// This slice of strings cant inherit from any interface, because of no type exists in this sample.
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)

	/*
		sort.StringSlice(s) es un Interface, a partir del cual ya poidemos invocar al método Sort()
	*/
	sort.StringSlice(s).Sort()
	//sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

	s2 := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(s2)
	fmt.Println(s2)

	// Sortering reverse
	s3 := []int{5, 2, 6, 3, 1, 4}              // unsorted
	sort.Sort(sort.Reverse(sort.IntSlice(s3))) // First we sort the slice and then we apply reverse function
	fmt.Println(s3)
}
