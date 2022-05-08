package main

import (
	"fmt"
	"strings"
)

func remove_whitespace(s []string) {
	for i := range s { // i, _ := range s {
		s[i] = strings.TrimSpace(s[i])
	}
}

func main() {

	// since slices are in one ways pointers to sections of arrays
	// slices can be passed to functions and arrays can be manipulated
	a := [...]string{"  a ", "  b", "c   "}
	s := a[:]
	fmt.Println(a, s)
	remove_whitespace(a[:])
	fmt.Println(a, s)
	// if you create additional slices, they still point to the exact same array

	// remember: the length of an array is part of its type
	// not the case with slices

	// in short: slices more versatile -> far more widely used than arrays
	// to create a slice directly

	s1 := []string{"a", "b", "c"}
	a1 := [...]string{"a", "b", "c"}
	fmt.Printf("%T %T\n", s1, a1) // []string [3]string
}
