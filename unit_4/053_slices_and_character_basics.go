package main

import "fmt"

func main() {

	array := [...]string{"a", "b", "c", "d", "e", "f"}
	// slicing is always from x (including) up to y (excluding)
	// in programming terminology: slicing is expressed with a half-open range
	fmt.Println(array[1:3]) // [b c]
	// slices are all views of the same array
	// you can still index into slices like arrays
	// assigning a new value to an element of a slice modifies the underlying array
	s := array[1:3]
	s[0] = "Z"
	fmt.Println(array) // [a Z c d e f]
	// ommitting first index defaults to beginning, ommitting last index defaults to end
	fmt.Println(array[:]) // [a Z c d e f]
	// slices can be thought of as pointers to arrays
	// https://blog.golang.org/slices

	// be aware that indices indicate the number of bytes
	s1 := `¿Cómo estás?`
	fmt.Println(s1[:6]) // prints ¿Cóm
	// see https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/
	// "In UTF-8, every code point from 0-127 is stored in a single byte.
	// Only code points 128 and above are stored using 2, 3, in fact, up to 6 bytes."
	// with standard characters of the ASCII set it works reliably
	s0 := "this :) is [][] a tes20t string with plain ASCII characters"
	fmt.Println(s0[:21])

}
