package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mobyDick, err := http.Get("http://wwww.gutenberg.org/files/2701/old/moby10b.txt")

	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(mobyDick.Body)
	// We close the buffer at the end of this execution using defer key word
	defer mobyDick.Body.Close()

	//Split by words
	scanner.Split(bufio.ScanWords)

	//Creating a slice to hold counts
	buckets := make([]int, 12)

	for scanner.Scan() {
		// We put the word on the map as key, with no value, no definition
		n := HashBucket(scanner.Text(), 12)
		buckets[n]++
	}

	fmt.Println(buckets[65:123])
}

func HashBucket(word string, buckets int) int {
	letter := int(word[0])
	return letter % buckets
}
