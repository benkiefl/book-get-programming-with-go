package main

import (
	"fmt"
	"sort"
)

func sortSlice(s []string, less func(i, j int) bool) {
	// if the less function passed is nil
	if less == nil {
		// define a default less function
		// in this case to sort alphabetically
		less = func(i, j int) bool { return s[i] < s[j] }
	}
	sort.Slice(s, less)
}

func main() {

	list := []string{"a", "c", "e", "b", "f", "h", "d", "g"}

	sortSlice(list, nil)

	fmt.Println(list)
}
