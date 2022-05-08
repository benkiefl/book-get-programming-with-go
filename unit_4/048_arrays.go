// hint: I did this unit twice before attempting
// the implementation of conway's game of life at end of chapter
// you might be smarter than me, friendly advice though:
// pay attention, take it slow, take it all in!
// arrays, slices, maps, sets, crucially important stuff to get right!

package main

import (
	"fmt"
)

func main() {
	// every element of an array has the same type
	var our_array [2]rune
	our_array[0] = 'a'
	our_array[1] = 'b'

	// [2]rune [3]rune [2]string [2]bool are all different types!
	// golang will catch a lot of potential misshaps this way

	// go can count for you, instead of planets := [3]string{}
	planets := [...]string{
		"Planet X",
		"Pluto",
		"Phibi", // notice the leading , it's needed
	}

	// you can do this in one line as well
	our_array_2 := [2]rune{'a', 'b'}

	fmt.Println("length of planets array is", len(planets))
	fmt.Println("length of our array is", len(our_array))
	fmt.Println("length of our array 2 is", len(our_array_2))
	// len() is a built in function
	// there aren't many: https://golang.org/pkg/builtin/

	// go does not detect every possible out of bounds error and may crash
	// the following doesn't give a warning but results in a panic: runtime error
	i := 3
	fmt.Println(our_array[i])
}
