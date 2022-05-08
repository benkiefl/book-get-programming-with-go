package main

import "fmt"

func main() {

	function()

}

func function() {

	// you can defer any function or method

	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")

	// defered statements are executed in LIFO (last in first out) order
	// think of each defered statement pushed onto a stack
	// and that stack then being worked through top to bottom

	// so the order will be: C - B - A

	return
}
