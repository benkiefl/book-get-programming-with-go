package main

import (
	"fmt"
	"math"
)

func main() {

	// let's assume a "rating" of 0.0 to 3.0
	// simplest: everything beginning with 0, 1, 2, 3
	data_raw := map[string]float64{
		"the plaza":    1.2,
		"ritz, paris":  .7,
		"claridge's":   2.3,
		"grand dame":   3.0,
		"raffles":      1.6,
		"ritz, london": 2.4,
	}

	data_sorted := make(map[float64][]string)

	// I like to call the variables k, v for key, value
	// same as I often use i, e for index, element (for arrays, slices)
	// not sure if this is good practice, for now it helps me
	for k, v := range data_raw {
		// math.Trunc() simply truncates floating point numbers
		// truncating is not rounding(!)
		v_truncated := math.Trunc(v)
		data_sorted[v_truncated] = append(data_sorted[v_truncated], k)
	}

	fmt.Println(data_sorted)
}
