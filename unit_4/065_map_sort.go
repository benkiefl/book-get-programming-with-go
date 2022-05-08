package main

import "fmt"

func main() {

	// think a rating of things from 1 to 3
	data_raw := map[string]int{
		"item_a": 3,
		"item_b": 1,
		"item_c": 2,
		"item_d": 3,
		"item_e": 3,
		"item_f": 2,
		"item_g": 1,
		"item_h": 1,
	}

	// let's sort it via a new map
	data_sorted := make(map[int][]string)

	// push data into the new data_sorted map
	for k, v := range data_raw {
		data_sorted[v] = append(data_sorted[v], k)
	}

	fmt.Println(data_sorted)

	// show all items rated three stars
	fmt.Println(data_sorted[3])
	// as stated before: map does not guarantee a specific order
	// run the program multiple times, results will be in various orders
}
