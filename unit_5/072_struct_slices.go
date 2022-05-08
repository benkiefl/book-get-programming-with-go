package main

import "fmt"

type location struct {
	name string
	lat  float64
	long float64
}

func main() {

	locations := []location{
		{name: "location 0", lat: 32.432, long: -21.3},
		{name: "location 1", lat: 22.1, long: 33.42},
		{name: "location 2", lat: 343.2, long: -170.32},
	}

	for i := range locations {
		fmt.Println(locations[i])
	}
}
