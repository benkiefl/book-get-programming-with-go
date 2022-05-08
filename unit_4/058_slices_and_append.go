package main

import "fmt"

func main() {

	sl := []string{"a", "b", "c"}

	sl1 := sl

	// since sl1 doesn't have enough room for append(sl, "d")
	// append allocates a new array with increased capacity
	sl1 = append(sl, "d")

	fmt.Printf("%v %v\n", len(sl), cap(sl))   // 3 3
	fmt.Printf("%v %v\n", len(sl1), cap(sl1)) // 4 6

	// since append has allocated a new array
	// modifying sl1 is different from modifying sl

	sl[1] = "Z"

	fmt.Println(sl, sl1)

}
