package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Many Go packages declare and export variables for errors

var (
	ErrBounds = errors.New("Out of Bounds.")
	ErrDigit  = errors.New("Wrong Digit.")
)

// if a function now returns an Error, the Caller can distinguish between those

func main() {

	rand.Seed(time.Now().UnixNano())

	err := randError()

	// this is just for demonstration
	// now when checking for an error you could for example do something like

	if err != nil {
		switch err {
		case ErrBounds, ErrDigit:
			fmt.Printf("Error Message: %v\n", err)
		default:
			fmt.Println(err)
		}
	}
	os.Exit(1)

}

// randomly returns an error value ; ErrBounds, ErrDigit or nil
func randError() error {
	randInt := rand.Intn(3)

	switch randInt {
	case 0:
		return ErrBounds
	case 1:
		return ErrDigit
	default:
		return nil
	}
}
