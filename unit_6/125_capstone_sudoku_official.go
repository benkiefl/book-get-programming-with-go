// this is 1:1 the solution from the book, only comments changed/added ; adding
// detailed comments is the best way of understanding foreign source code

package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	rows, columns = 9, 9
	empty         = 0
)

// Cell is a square on the Sudoku grid
// additional bool to each field to keep track of the init configuration
// I did that with a second matrix IP stored the init config in and compared to
type Cell struct {
	digit int8
	fixed bool
}

// Grid is a Sudoku grid, every value is of type Cell (struct with int8 and bool variables)
type Grid [rows][columns]Cell

// Errors that could occur
var (
	ErrBounds     = errors.New("out of bounds")
	ErrDigit      = errors.New("invalid digit")
	ErrInRow      = errors.New("digit already present in this row")
	ErrInColumn   = errors.New("digit already present in this column")
	ErrInRegion   = errors.New("digit already present in this region")
	ErrFixedDigit = errors.New("initial digits cannot be overwritten")
)

// NewSudoku is a constructor function, makes a new Sudoku grid
// takes a complete Matrix of [][]int8
func NewSudoku(digits [rows][columns]int8) *Grid {
	// declare a Grid (we will it return via &Grid)
	var grid Grid
	// ranges over the matrix the function was passed
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			// stores value digits[r][c] into d
			d := digits[r][c]
			// only if d isn't 0 it will be stored into the grid, and bool
			// "flipped to true" (will be 0 and false by default anyways)
			if d != empty {
				grid[r][c].digit = d
				grid[r][c].fixed = true
			}
		}
	}
	// return the grid
	return &grid
}

// Set a digit on a Sudoku grid
func (g *Grid) Set(row, column int, digit int8) error {
	// this looks very tidy, I have to admit
	switch {
	case !inBounds(row, column):
		return ErrBounds
	case !validDigit(digit):
		return ErrDigit
	// simply returns the .fixed Bool of our Cell structure
	case g.isFixed(row, column):
		return ErrFixedDigit
	// checks if value is already in the row
	case g.inRow(row, digit):
		return ErrInRow
	// checks if value is already in volumn
	case g.inColumn(column, digit):
		return ErrInColumn
	// checks if value is already in region (what I called squares)
	case g.inRegion(row, column, digit):
		return ErrInRegion
	}
	// if all tests are passed, store the digit in the grid and return nil
	g[row][column].digit = digit
	return nil
}

// Clear clears a single field, does !inBounds and g.isFixed tests, then zeroes the field
func (g *Grid) Clear(row, column int) error {
	switch {
	case !inBounds(row, column):
		return ErrBounds
	case g.isFixed(row, column):
		return ErrFixedDigit
	}
	g[row][column].digit = empty
	return nil
}

// inBounds checks if values are in bounds (based on the constants rows, columns)
func inBounds(row, column int) bool {
	if row < 0 || row >= rows || column < 0 || column >= columns {
		return false
	}
	return true
}

// validDigit checks if digit is between 1 and 9
func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

// inRow checks if digit already appears in the row
func (g *Grid) inRow(row int, digit int8) bool {
	for c := 0; c < columns; c++ {
		if g[row][c].digit == digit {
			return true
		}
	}
	return false
}

// inColumn checks if digit already appears in the column
func (g *Grid) inColumn(column int, digit int8) bool {
	for r := 0; r < rows; r++ {
		if g[r][column].digit == digit {
			return true
		}
	}
	return false
}

// inRegion checks if a given digit is already in the region / square
func (g *Grid) inRegion(row, column int, digit int8) bool {
	// row, column are integer values, will be truncated upon division
	// this is the elegant solution I've been missing
	// if you know where to "start" you can simply check a 3x3 matrix from there
	startRow, startColumn := row/3*3, column/3*3
	// check till startRow+3
	for r := startRow; r < startRow+3; r++ {
		// and till startColumn+3
		for c := startColumn; c < startColumn+3; c++ {
			// and then it's simply comparing
			if g[r][c].digit == digit {
				return true
			}
		}
	}
	return false
}

// isFixed checks bool value (.fixed in the struct) and returns it
func (g *Grid) isFixed(row, column int) bool {
	return g[row][column].fixed
}

func main() {
	// use the Constructur function, remember our Sudoku has struct values with
	// bools to see which values have been fixed from the start, so it's not
	// simply returning the same matrix or something
	s := NewSudoku([rows][columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	err := s.Set(1, 1, 4)
	if err != nil {
		fmt.Println(err)
		// exits upon error
		os.Exit(1)
	}

	// printing the rows allows us to see the sudoku's structure
	// it contains structs (which consist of an int8 and a bool)
	for _, row := range s {
		fmt.Println(row)
	}

}
