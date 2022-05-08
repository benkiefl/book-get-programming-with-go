// I thought I would come up with my own example

// the technique of using an "assembly line" of functions using channels is
// called pipeline ; it allows for processing of large amounts of data without
// large amount of memory ; it also is a model to think about problems, break
// them down into sub-processes ; might help to find easy solutions
// you might appreciate https://www.youtube.com/watch?v=oV9rvDllKEg

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go reader(c0)
	go filter(c0, c1)
	printer(c1)

}

// reader function reads from a file, converts data into a []string of words,
// split on " " (spaces), and sends the data onto a channel
func reader(downstream chan string) {

	data, err := ioutil.ReadFile("137_text_with_expletives.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	// my data is in bytes, convert it to string
	text := string(data)
	// split it on " " and store it in []string
	words := strings.Split(text, " ")

	// range over []string and send the words downstream
	for _, word := range words {
		downstream <- word
	}
	// when finished send an empty string
	downstream <- ""
	// this is called a sentinel value
	// https://en.wikipedia.org/wiki/Sentinel_value
}

func filter(upstream, downstream chan string) {
	var word string
	// endless for loop
	for {
		// take word from upstream
		word = <-upstream
		// check if word is an expletive (very simplistic here of course)
		if word != "fuck" && word != "shitty" && word != "fucking" {
			// if it isn't, send it downstream
			downstream <- word
		}
		// if word is an empty string, return
		if word == "" {
			return
		}
	}
}

func printer(upstream chan string) {
	// create a file
	f, err := os.Create("137_text_without_expletives.txt")
	// before you return, close the file
	defer f.Close()

	if err != nil {
		fmt.Println("Error upon file creation", err)
		return
	}

	// endless loop
	for {
		// take word from upstream
		word := <-upstream
		// if empty string return
		if word == "" {
			return
		}
		// print the word to the file, separate using " "
		fmt.Fprint(f, word, " ")
	}
}
