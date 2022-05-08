package main

import (
	"fmt"
	"math/big"
)

func main() {

	const LIGHTSPEED = 299792 // km/s
	const SECONDS_PER_DAY = 24 * 60 * 60
	const DISTANCE_ALPHA = 41.3e12
	const DISTANCE_ANDROMEDA = 24e18

	fmt.Printf("alpha centauri is %d km away\n", int(DISTANCE_ALPHA))
	travel_time := DISTANCE_ALPHA / LIGHTSPEED / SECONDS_PER_DAY
	fmt.Printf("traveling at lightspeed it takes %d days\n", int(travel_time))

	// big.Ints are troublesome to work with and slower
	// but this way you can calculate with numbers of (theoretically) any size

	distance_andromeda := new(big.Int)
	distance_andromeda.SetString("24000000000000000000", 10)
	// since even a uint64 can't hold the above value
	// you have to assign it to our big.Int variable via conversion from string
	lightspeed := big.NewInt(LIGHTSPEED)
	seconds_per_day := big.NewInt(SECONDS_PER_DAY)

	seconds := new(big.Int)
	days := new(big.Int)
	years := new(big.Int)
	// all variables involved have to be big.Ints

	fmt.Printf("andromeda is %d km away\n", distance_andromeda)
	// even normal division doesn't work, you have to use the .Div method (package math/big)
	seconds.Div(distance_andromeda, lightspeed)
	days.Div(seconds, seconds_per_day)
	years.Div(days, big.NewInt(365))
	fmt.Printf("traveling at lightspeed it takes %d days ; that's %d years\n", days, years)

	// due to untyped constants able to hold values of any size, we don't need big.Ints
	// behind the scenes untyped consts are handled with help of the big package
	// constants and big.Ints aren't interchangeable
	// for example untyped constants can easily overflow your ints in use
	fmt.Printf("andromeda is %d km away\n", int(DISTANCE_ALPHA))
	travel_time = DISTANCE_ANDROMEDA / LIGHTSPEED / SECONDS_PER_DAY
	fmt.Printf("traveling at lightspeed it takes %d days ; that's %d years\n", int(travel_time), years)
}
