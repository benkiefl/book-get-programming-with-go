package main

import "fmt"

// lesson 17 describes slices as windows into arrays
// sounds a lot like pointers, doesn't it?
// because it is, see https://blog.golang.org/slices-intro#TOC_4.

func main() {

	// there is still a use cases for explicitly passing slices around as pointers
	// and that is when modifying the slice itself, meaning its length, capacity or starting offset

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("%d %d %T %v\n", len(nums), cap(nums), nums, nums)

	reclass(nums)
	fmt.Printf("%d %d %T %v\n", len(nums), cap(nums), nums, nums)

	reclassify(&nums)
	fmt.Printf("%d %d %T %v\n", len(nums), cap(nums), nums, nums)

	nums = RECLASSIFY(nums)
	fmt.Printf("%d %d %T %v\n", len(nums), cap(nums), nums, nums)
}

// this won't work as intended
func reclass(n []int) {
	n = n[0:4]
}

func reclassify(n *[]int) {
	*n = (*n)[0:4]
}

// probably an even more elegant solution
func RECLASSIFY(n []int) []int {
	return n[0:2]
}
