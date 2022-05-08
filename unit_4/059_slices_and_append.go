package main

import "fmt"

func main() {

	s0 := []string{"a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j", "k", "l", "m"}
	s1 := s0

	s1 = append(s1, "n", "o")

	fmt.Printf("%v %v\n", len(s0), cap(s0)) // 13 13
	fmt.Printf("%v %v\n", len(s1), cap(s1)) // 15 26

	// it will be doubled up to a certain limit (if cap(slice) < 1024)
	// from there on other heuristics are used
	// https://github.com/golang/go/blob/master/src/runtime/slice.go#L115
	// good idea to get to know the golang github repos
	// https://github.com/golang

}
