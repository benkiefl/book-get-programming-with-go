// https://en.wikipedia.org/wiki/Set_(abstract_data_type)
// there is no extra type for sets, but you can use maps
// in case you're some super math nerd
// https://www.cl.cam.ac.uk/~gw104/STfCS2010.pdf

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var values []int
	max_values := 10

	for i := 0; i < max_values; i++ {
		values = append(values, rand.Intn(11))
	}

	set := make(map[int]bool)

	for _, e := range values {
		set[e] = true
	}

	fmt.Println(values)
	// since set is a map[int]bool
	// it's easy to use our set in checks
	if set[0] {
		fmt.Println("0 is in the set")
	}

	// maps can't be relied upon to have a specific order
	// in order to sort a map, you can convert it to a slice

	// slice_of_set, slice []int with 0 length, and len(set) as cap
	slice_of_set := make([]int, 0, len(set))

	// remember: it's just a map, there's no extra data type for a set
	for k, _ := range set {
		slice_of_set = append(slice_of_set, k)
	}

	sort.Ints(slice_of_set)

	fmt.Println(slice_of_set)

	/*	I figure slices and maps will be the most important types
		to work with various data collections, do data analysis in general
		quite flexible and powerful
	*/
}
