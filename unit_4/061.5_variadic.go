// this is about something I found peculiar
// feel free to ignore

package main

import (
	"fmt"
	"strings"
)

func variadic(test ...string) {

	for _, e := range test {
		fmt.Print(strings.ToUpper(e), " ")
	}
	fmt.Println()

}

func main() {

	str := "test"
	sl := []string{"and this", "should work", "as well"}

	variadic(append([]string{str}, sl...)...)
	// but why doesn't variadic(str, sl...) work?
	// cannot use sl (type []string) as type string in argument to variadic

	// so in above line we're basically doing append()
	// but append needs a []string as first argument, you can't append a slice to a string
	// step 0 : convert string str to []string by doing []string{str}
	// step 1 : append []string{str} and sl...
	// step 3 : now pass the the newly appended string as one to the function variadic by using ...

	/*
		it's been pointed out to me that such a thing should probably not come up in go
		in other words: needing to insert in the beginning of the slice is probably
		indicative of some fundamental design flaw in the program to begin with
		I found above solution interesting enough to make a note of it though
	*/

}
