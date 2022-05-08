package main

import "fmt"

func main() {

	// a slice declared without a composite literal or make will have value of nil
	var slice0 []string
	// otherwise the slice is empty but not nil
	slice1 := []string{}
	slice2 := make([]string, 0, 0)
	// let's see
	fmt.Println(slice0 == nil, slice1 == nil, slice2 == nil)

	// however: range, len, append all work perfectly fine with nil slices
	var items []string
	fmt.Println("length:", len(items))
	// as you can see, range with an empty or nil slice doesn't do anything
	i := 5
	for i = range items {
	}
	fmt.Println("range doesn't execute, i is still", i) // i still 5

	items = append(items, "item_a", "item_b")
	fmt.Println("appended items, items is now", items)
}
