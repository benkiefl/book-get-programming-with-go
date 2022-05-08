package main

import (
	"fmt"
	"strings"
)

func prepend(word ...string) {
	// creates a word parameter as a slice of strings
	// it has to be dealt with as a slice
	for i := range word {
		fmt.Println(strings.ToUpper(word[i]))
	}
}

func main() {

	str_0 := "a"
	str_1 := "b"
	str_2 := "c"

	prepend(str_0, str_1, str_2)

	// you can pass a slice as single arguments by using ... as well

	sl := []string{"a", "b", "c"}

	prepend(sl...)

	// prepend(sl) won'twork here, prepend() expects number of single arguments

}
