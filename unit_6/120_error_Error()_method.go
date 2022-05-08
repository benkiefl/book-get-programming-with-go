package main

import (
	"errors"
	"fmt"
	"strings"
)

const rows, cols = 9, 9

var (
	ErrBounds = errors.New("Out of Bounds.")
	ErrDigit  = errors.New("Wrong Digit.")
)

type Grid [rows][cols]int8

// type SudokuError will satisfy the Error() interface
type SudokuError []error

// SudokuError will be of type [] error ; in Error() we will join the (possibly
// multiple) error message together and return them as one string
func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		// err.Error() converts error to a string
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

// Set method sets a value for a Grid, performs various Error checks, returns
// nil upon successful completion
func (g *Grid) Set(row, col int, digit int8) error {
	// about above line: always return errors as type error
	// not as SudokuError et cetera
	var errs SudokuError
	if !inBounds(row, col) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}
	g[row][col] = digit
	return nil
}

// the Print method prints a given Grid
func (g *Grid) Print() {
	for i := range g {
		for l := range g[i] {
			fmt.Printf("%v ", g[i][l])
		}
		fmt.Println()
	}
}

func main() {

	var Sudoku Grid

	err := Sudoku.Set(10, 0, 10)
	// Sudoku.Set() returns a value of type error
	// so how to access the individual erros?
	// using type assertion https://tour.golang.org/methods/15
	if err != nil {
		// if the err interface has a type SudokuError, ok will be true
		// errs will be our error Slice
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occured:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
	}
	// err.(SudokuError) attempts to convert the err value from the error
	// interface tyupe to the concrete SudokuError type

}

// inBounds checks whether row, col are in Bounds of our Sudoku Matrix
func inBounds(row, col int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if col < 0 || col >= cols {
		return false
	}
	return true
}

// validDigit checks if our value is a legal value for Sudoku
func validDigit(n int8) bool {
	if n >= 1 && n <= 9 {
		return true
	}
	return false
}
