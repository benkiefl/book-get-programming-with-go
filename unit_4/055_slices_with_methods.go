package main

import (
	"fmt"
	"sort"
)

func main() {

	// there are certain methods that work only on slices
	// adds to slices' versatility (especially over arrays)

	s := []string{"c", "b", "e", "d", "a"}

	fmt.Println(s)
	sort.StringSlice(s).Sort()
	fmt.Println(s)
	// the sort package (see import above) declares a StringSlice type
	// attached to the StringSlice type is a Sort() method

	// the sort package has a Strings helper function
	// it converts the type conversion and calls the Sort method
	// planets := [...]string{"c", "b", "e", "d", "a"}
	// sort.Strings(planets)
	// planets has to be a slice though

	planets1 := []string{"c", "b", "e", "d", "a"}
	fmt.Println(planets1)
	sort.Strings(planets1)
	fmt.Println(planets1)
}
