// popular interfaces in the stdlib are
// fmt.Stringer, io.Reader, io.Writer, json.Marshaler

package main

import "fmt"

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long coordinate
}

// the Stringer interface is part of fmt package, it's defined as follows
// type Stringer interface {
//     String() string
// }
// Stringer is implemented by any value that has a String method
// see https://pkg.go.dev/fmt?utm_source=gopls#Stringer

func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'%v\" %c", c.d, c.m, c.s, c.h)
}

func main() {

	elysium := location{
		lat:  coordinate{d: 4, m: 30, s: 0.0, h: 'N'},
		long: coordinate{d: 135, m: 54, s: 0.0, h: 'E'},
	}

	fmt.Println("Elysium is at", elysium)

}

// summary
// interface types specify required behaviors with a set of methods
// interfaces are satisfied implicity by new or existing code in any package
// a structure will satisfy the interfaces that embedded types satisfy
// keep interfaces small (follow example set by interfaces in the standard library)
