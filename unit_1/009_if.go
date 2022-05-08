package main

import "fmt"

func main() {

	var command string = "stay"

	// important: the first if which is true will be executed
	// other comparisons further down won't even be evaluated

	if command == "go east" {
		fmt.Println("You head further east")
	} else if command == "stay" {
		fmt.Println("You stay, awaiting almost certain death")
	} else {
		fmt.Println("Invalid Input entered")
	}

}
