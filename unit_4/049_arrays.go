package main

import "fmt"

func main() {

	planets_dwarf := [...]string{
		"Ceres",
		"Pluto",
		"Haumea",
		"Makemake",
		"Eris",
	}

	// iterating over an array
	for i := 0; i < len(planets_dwarf); i++ {
		fmt.Println(i, planets_dwarf[i])
	}

	// alternative: using range, less code, less chances for mistakes
	// range returns two values, the index, and copy of the element
	for index, element := range planets_dwarf {
		fmt.Println(index, element)
	}

	// you can use underscore if you only need one or the other
	for _, element := range planets_dwarf {
		fmt.Println(element)
	}

}
