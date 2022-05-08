package main

import (
	"errors"
	"fmt"
	"os"
)

const rows, cols = 9, 9

type Grid [rows][cols]int8

// validating parameters at the beginning of a method guards the remainder of
// the method from worrying about bad input

func (g *Grid) Set(row, col int, digit int8) error {
	// we've put the logic checking if row, col are legal inputs into its own function
	if !inBounds(row, col) {
		// errors.New returns an error which formats as the given text
		return errors.New("Out of Bounds")
		// it's basically a constructor function for errors
	}

	g[row][col] = digit
	return nil
}

// inBounds checks if row, col are legal inputs
// to initialize a value in a Sudoku Grid ; returns false it not, true if yes
func inBounds(row, col int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if col < 0 || col >= cols {
		return false
	}
	return true
}

func main() {

	var g Grid

	err := g.Set(10, 0, 5)
	if err != nil {
		fmt.Printf("An error occured: %v.\n", err)
		os.Exit(1)
	}
}

// let err be whatever is informative, as soon as err is _something_: err !=nil succeeds
// if !inBounds(row, col) {} : pretty neat ; possible since inBounds() returns a Bool value
// don't be afraid to outsource parts of error handling
// last time: our own safeWriter struct which had a method writeln, error handling built into it
// this time: inBounds() which returns a bool and upon failure creates a new error via errors.New()
