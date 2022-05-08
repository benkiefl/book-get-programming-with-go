// I added reading from a file, and split using strings.Split instead of
// strings.Fields ; this program all in all is very similar to
// 137_goroutines_expletives_filtering.go with the detection logic taken from
// 142_exercise_remove_identical.go

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	c0 := make(chan string)
	c1 := make(chan string)
	go source(c0)
	go filter(c0, c1)
	print(c1)
}

func source(down chan string) {
	data, err := ioutil.ReadFile("142_text_with_duplicates.txt")
	if err != nil {
		fmt.Println("error while reading from file:", err)
	}
	text := string(data)
	words := strings.Split(text, " ")
	for _, word := range words {
		down <- strings.TrimSuffix(word, "\n") // to remove newlines
	}
	close(down)
}

func filter(up, down chan string) {
	prev := ""
	for elem := range up {
		if prev != elem {
			down <- elem
			prev = elem
		}
	}
	close(down)
}

func print(up chan string) {
	for elem := range up {
		fmt.Print(elem, " ")
	}
}
