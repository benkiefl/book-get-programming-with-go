package main

import "fmt"

func main() {

	planets := []string{
		"a", "b", "c", "d",
		"e", "f", "g", "h"}

	fmt.Println("original array:", planets)

	sl_0 := planets[:]
	sl_1 := planets[0:4:4]
	sl_2 := planets[0:4]

	fmt.Printf("%v %v %v\n", cap(sl_0), cap(sl_1), cap(sl_2)) // 8 4 8

	// what's the real difference here? that's a little bit tricky

	sl_1 = append(sl_1, "Z") // has cap of 4, appending outgrows cap, new array allocated
	// so above line does _not_ modify planets, let's check
	fmt.Println(planets)

	sl_2 = append(sl_2, "Z") // has cap of 8, appending doesn't outgrow cap
	// above line modifies planets, in fact it modifies planets[4] (e)
	fmt.Println(planets)

	fmt.Println(sl_1, sl_2)                     // [a b c d Z] [a b c d Z]
	fmt.Printf("%v %v\n", cap(sl_1), cap(sl_2)) // 8 8

	// you can avoid extra new allocations by preallocating a slice
	// use builtin make function

	planets_2 := make([]string, 0, 10) // length 0 and capacity 10
	// the capacity argument is optional
	planets_3 := make([]string, 10) // length and capctiy 10
	planets_2 = append(planets_2, "a", "b", "c")
	planets_3 = append(planets_3, "a", "b", "c")
	fmt.Printf("%v %v %v\n", planets_2, len(planets_2), cap(planets_2))
	fmt.Printf("%v %v %v\n", planets_3, len(planets_3), cap(planets_3))

	// notice: leading empty elements are shown simply as whitespaces!
}
