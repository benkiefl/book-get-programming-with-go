package main

import "fmt"

func main() {

	var distance_min int = 56000000
	var travel_time int = 28 * 24 // hours

	fmt.Println("travel speed would have to be", distance_min/travel_time, "km/h")

}
