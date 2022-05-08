package main

import "fmt"

func main() {

	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	frequency := make(map[float64]int)

	// with maps, range returns key, value
	for _, t := range temperatures {
		frequency[t]++
	}

	for t, num := range frequency {
		fmt.Printf("%+.2f occurs %d times\n", t, num)
	}

	// careful: maps are not guaranteed to be in a certain order

}
