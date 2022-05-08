package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Grid [9][9]int

// our Control Matrix to store the Init configuration
var control Grid

var (
	ErrInit = errors.New("Field was set during Init and can't be changed")
	ErrRule = errors.New("Not a legal entry")
)

// Set method sets a certain value, also can overwrite a value (unless its
// value was set during Init)
func (g *Grid) Set(row, col int, val int) error {
	if control[row][col] != 0 {
		return ErrInit
	}
	if check_legal(g, row, col, val) {
		g[row][col] = val
	} else {
		return ErrRule
	}

	return nil
}

func (g *Grid) Init() {

	g.Wipe()

	rand.Seed(time.Now().UnixNano())

	for i := range g {
		for l := range g[i] {
			// by creating one of 24 possible values, but only using it if
			// between 1 and 9 I attribute a chance to a field being initialized
			rval := rand.Intn(24) + 1
		AGAIN:
			if rval >= 1 && rval <= 9 {
				if check_legal(g, i, l, rval) {
					g[i][l] = rval
					control[i][l] = rval
				} else {
					// if initialization fails due to legality of the value, we
					// still want the field initialized, so: new random value
					// between 1 and 9 and repeat until we find a valid one
					rval = rand.Intn(9) + 1
					goto AGAIN
				}
			}
		}
	}
	// why 24/9? because 24/9 means each field has a chance of 37.5% to be
	// initialized ; the example Sudoku in the book has 30/81 fields initialized
	// (which is 30/81, so ~37.04%), close enough
}

func (g *Grid) Wipe() {
	for i := range g {
		for l := range g[i] {
			g[i][l] = 0
			control[i][l] = 0
		}
	}
}

func check_legal(g *Grid, row, col int, val int) bool {

	for i := range g[row] {
		if val == g[row][i] {
			return false
		}
	}

	for i := range g {
		if val == g[i][col] {
			return false
		}
	}

	// slices sq0 to sq8 are the squares of the Sudoku
	sq0 := [3][]int{g[0][0:3], g[1][0:3], g[2][0:3]}
	sq1 := [3][]int{g[0][3:6], g[1][3:6], g[2][3:6]}
	sq2 := [3][]int{g[0][6:9], g[1][6:9], g[2][6:9]}
	sq3 := [3][]int{g[3][0:3], g[4][0:3], g[5][0:3]}
	sq4 := [3][]int{g[3][3:6], g[4][3:6], g[5][3:6]}
	sq5 := [3][]int{g[3][6:9], g[4][6:9], g[5][6:9]}
	sq6 := [3][]int{g[6][0:3], g[7][3:6], g[8][6:9]}
	sq7 := [3][]int{g[6][0:3], g[7][3:6], g[8][6:9]}
	sq8 := [3][]int{g[6][0:3], g[7][3:6], g[8][6:9]}
	sq := sq0

	// so we don't have to repeat writing code ranging over the square, check
	// which square we are in and assign it to sq
	if row < 3 && col < 3 {
		sq = sq0
	} else if row < 3 && col < 6 {
		sq = sq1
	} else if row < 3 && col < 9 {
		sq = sq2
	} else if row < 6 && col < 3 {
		sq = sq3
	} else if row < 6 && col < 6 {
		sq = sq4
	} else if row < 6 && col < 9 {
		sq = sq5
	} else if row < 9 && col < 3 {
		sq = sq6
	} else if row < 9 && col < 6 {
		sq = sq7
	} else if row < 9 && col < 9 {
		sq = sq8
	}

	// now we write the logic once
	for i := range sq {
		for l := range sq[i] {
			if sq[i][l] == val {
				return false
			}
		}
	}

	// if all our checks fails
	// not in row, not in column, not in square
	// then return true

	return true
}

func (g *Grid) Reset() {
	for i := range g {
		for l := range g[i] {
			g[i][l] = control[i][l]
		}
	}
}

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
	var r, c, v int
	var choice string

	fmt.Printf("(I)nit, (S)et, (W)ipe, (R)eset, (C)lear Screen, (P)rint, (E)xit : ")

	Sudoku.Init()

	for {
		fmt.Scanln(&choice)
		switch choice {
		case "i", "I":
			Sudoku.Init()
		case "s", "S":
			fmt.Printf("Enter row: ")
			fmt.Scanln(&r)
			fmt.Printf("Enter column: ")
			fmt.Scanln(&c)
			fmt.Printf("Enter value: ")
			fmt.Scanln(&v)
			if err := Sudoku.Set(r, c, v); err != nil {
				fmt.Println(err)
			}
		case "w", "W":
			Sudoku.Wipe()
		case "r", "R":
			Sudoku.Reset()
		case "c", "C":
			clear_screen()
			fmt.Printf("(I)nit, (S)et, (W)ipe, (R)eset, (C)lear Screen, (P)rint, (E)xit: ")
		case "p", "P":
			Sudoku.Print()
		case "e", "E":
			os.Exit(1)
		}
	}

}

func clear_screen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
