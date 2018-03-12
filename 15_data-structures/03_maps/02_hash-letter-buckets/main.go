package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	//HASH TABLES. Go provides a built-in map type that implements a hash table.
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")

	if err != nil {
		log.Fatalln(err)
	}

	// We can read all the request body and print this information. Obtain []bytes
	bs, _ := ioutil.ReadAll(res.Body)
	str := string(bs)
	fmt.Println(str)

	// We can create a map to store the information
	words := make(map[string]string)

	// NewScanner returns a new Scanner to read from r. Receives a reader.
	sc := bufio.NewScanner(res.Body)

	// We close the buffer at the end of this execution using defer key word
	defer res.Body.Close()

	//Split by words
	sc.Split(bufio.ScanWords)

	// While there are words, we iterate
	for sc.Scan() {
		// We put the word on the map as key, with no value, no definition
		words[sc.Text()] = ""
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 0
	for k, _ := range words {
		fmt.Println(k)

		if i == 200 {
			break
		}
		i++
	}

	mobyDick, err := http.Get("http://wwww.gutenberg.org/files/2701/old/moby10b.txt")

	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(mobyDick.Body)
	// We close the buffer at the end of this execution using defer key word
	defer mobyDick.Body.Close()

	//Creating a slice to hold counts
	buckets := make([]int, 200)

	for scanner.Scan() {
		// We put the word on the map as key, with no value, no definition
		n := HashBucket(scanner.Text())
		buckets[n] += 1
	}

	fmt.Println(buckets[65:123])

}

func HashBucket(word string) int {
	return int(word[0])
}
