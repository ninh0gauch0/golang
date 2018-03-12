package main

import (
	"encoding/json"
	"fmt"
)

//Person struct with tags
type Person struct {
	First       string
	Last        string `json:"-"`
	Age         int    `json:"wisdom score"`
	notExported int    // This variable is unexported, it's not be printed
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
