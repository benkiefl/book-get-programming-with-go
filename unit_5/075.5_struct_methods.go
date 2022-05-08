// tried the conversion the other way around
// and wanted to utilize a struct as well
// ended up implementing quite a lot of poor workarounds
// and implementing stuff not yet taught in the book
// PS: but it works

package main

import (
	"fmt"
	"math"
)

type CoordinateDMS struct {
	d, m, s float64
	h       rune
}

type CoordinateDegrees struct {
	degrees float64
}

func (c CoordinateDMS) degree() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1.0
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// method dms on CoordinateDegrees struct converts and returns CoordinateDMS struct
func (c CoordinateDegrees) dms(lat_or_long string) CoordinateDMS {
	var hemisphere rune
	// conditional based on a string that has to be passed along
	// even to my untrained eyes this is extremely shady
	if lat_or_long == "lat" {
		if c.degrees < 0 {
			hemisphere = 'S'
		} else {
			hemisphere = 'N'
		}
	}
	if lat_or_long == "long" {
		if c.degrees < 0 {
			hemisphere = 'W'
		} else {
			hemisphere = 'E'
		}
	}

	// did the steps first on paper and then just implemented them in an ugly manner
	c.degrees = math.Abs(c.degrees)
	degrees := math.Floor(c.degrees)
	minutes := math.Floor((c.degrees - degrees) * 60)
	minutes_raw := (c.degrees - degrees) * 60
	seconds := (minutes_raw - minutes) * 60
	return_coordinates := CoordinateDMS{d: degrees, m: minutes, s: seconds, h: rune(hemisphere)}
	return return_coordinates
}

// https://golang.org/pkg/fmt/#GoStringer
// https://tour.golang.org/methods/17
func (s CoordinateDMS) String() string {
	return fmt.Sprintf("%.0f°%.0f'%.2f\" %q", s.d, s.m, s.s, s.h)
}

func main() {

	lat := CoordinateDMS{4, 35, 22.2, 'S'}
	long := CoordinateDMS{137, 26, 30.12, 'E'}

	fmt.Println(lat.degree(), long.degree())

	lat_2 := CoordinateDegrees{-4.5895}
	long_2 := CoordinateDegrees{137.4417}

	fmt.Println(lat_2.dms("lat"), long_2.dms("long"))

}
