package main

import "fmt"

func main() {

	c0 := make(chan string)
	c1 := make(chan string)
	go source(c0)
	go filter(c0, c1)
	print(c1)
}

func source(down chan string) {
	sentences := []string{"a", "b", "b", "c", "c", "c", "d", "e"}
	for _, sentence := range sentences {
		down <- sentence
	}
	close(down)
}

func filter(up, down chan string) {
	prev := ""
	for sentence := range up {
		if sentence != prev {
			down <- sentence
			prev = sentence
		}
	}
	close(down)
}

func print(up chan string) {
	for element := range up {
		fmt.Println(element)
	}
}
