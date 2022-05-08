package main

import "fmt"

func main() {

	const LIGHTSPEED = 299792 // km/s

	// in km:
	var distance_short = 56000000
	var distance_long = 401000000

	fmt.Println(distance_short/LIGHTSPEED, "seconds")
	fmt.Println(distance_long/LIGHTSPEED, "seconds")

}
