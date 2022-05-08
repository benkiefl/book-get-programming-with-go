package main

import "fmt"

func main() {

	const HOURS_PER_DAY = 24

	var (
		distance = 96300000 // km
		speed    = 100800   // km/h
	)

	// var distance, speed = 96300000. 100800
	// though consider readability

	fmt.Println(distance/speed/HOURS_PER_DAY, "days")

}
