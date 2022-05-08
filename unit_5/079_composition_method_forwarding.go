// in OOP objects are composed of smaller objects
// is called object composition / composition

// in Go this is achieved with structures and _embeddindg_ to forward methods

package main

import "fmt"

// poorly_structured weather report struct
type poor_report struct {
	sol       int
	high, low float64 // high low what?
	lat, long float64 // what is lat long?
}

// properly structured weather report struct (object composition)
type report struct {
	sol         int
	temperature temperature // structures within make things clear
	location    location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

// method forwarding
func (r report) average() celsius {
	return r.temperature.average()
}

type celsius float64

func main() {

	report := report{
		sol:         15,
		temperature: temperature{high: 23, low: -10},
		location:    location{lat: 59.94885521242605, long: 30.31703316085591},
	}

	fmt.Printf("%+v\n", report)

	fmt.Printf("%+v\n", report.temperature.average())

	// due to the method forwarding above, we can access average() directly via report
	fmt.Printf("%+v\n", report.average())
	// method forwarding is very practical but adds a lot of clutter if done like here
	// see next file
}
