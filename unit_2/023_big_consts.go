package main

import "fmt"

func main() {

	const LIGHTSPEED = 299792
	const SECONDS_PER_DAY = 24 * 60 * 60
	const CANIS = 236e18

	fmt.Printf("the canis major dwarf galaxy is %.0f km away\n", CANIS)
	days := CANIS / LIGHTSPEED / SECONDS_PER_DAY
	years := days / 365.25
	fmt.Printf("travel time with lightspeed %.0f days or %.0f years\n", days, years)

}
