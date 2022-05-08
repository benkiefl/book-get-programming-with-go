// uninitialized variables have the zero value corresponding to their type
// for pointers, maps, slices, interfaces the zero value is nil
// string it's an empty string, int 0, float 0.0, etc

package main

import "fmt"

func main() {

	var pointer *int

	fmt.Println(pointer)

	// dereferencing a nil pointer usually crashes the program
	// fmt.Println(*pointer)
	// panic: runtime error: invalid memory address or nil pointer dereference

	// it's simple to guard against this

	if pointer != nil {
		fmt.Println(*pointer)
	}

	// nil pointers can be useful but also be a burden
	// nil pointers in Go are less prevalent than null pointers in some other languages
}
