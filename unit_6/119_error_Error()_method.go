// see https://tour.golang.org/methods/19

package main

import (
	"fmt"
	"os"
	"time"
)

// there's a built in interface called error
// a type satisfies the interface by implementing Error()
// it can then be returned as an error type

// by convention the names of custom error types end with Error
// in some cases they are just called Error, like url.Error from the url package
type MyError struct {
	When time.Time
	What string
}

// a method Error() on *MyError means we can return the struct as an error
func (m *MyError) Error() string {
	return fmt.Sprintf("%v happened at %v", m.What, m.When)
}

// you can use a blank identifier to perform a type check of sorts
// if (*MyError)(nil) isn't of type error this will fail during build
// this way you could check if a certain type implements the error interface
var _ error = (*MyError)(nil)

func main() {

	value := 5

	if err := quadruple(&value); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	value = -5

	if err := quadruple(&value); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// quadruples an integer value, only works on integers > 0
func quadruple(val *int) error {

	if *val <= 0 {
		return &MyError{time.Now(), "Illegal Value"}
	}

	*val *= 4
	return nil
}
