package main

import "fmt"

func main() {

	var count int = 0

	for count < 5 {
		fmt.Println(count)
		count++
	}

	// more concise
	// c-like for loop

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// break can be used to exit a loop at any time

	i := 0

	for { // for {} without any condition -> infinite loop
		fmt.Println(i)
		if i == 3 {
			break // jumps out of loop
		}
		i++
	}
}
