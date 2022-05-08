/* Conway's Game of Life
1. Any live cell with two or three live neighbours survives.
2. Any dead cell with three live neighbours becomes a live cell.
3. All other live cells die in the next generation. Similarly, all other dead cells stay dead.
https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
*/

/*
can recommend Golly (cross-platform, even on Android)
*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	width       = 160
	height      = 50
	generations = 1000
)

type Universe [][]bool

// creates and returns slice of type Universe using constants width and height
func NewUniverse() Universe {
	// create [][]bool slice of constant _height_ slices
	new_u := make(Universe, height)
	// use range to pre-allocate every slice with constant _width_
	for i := range new_u {
		new_u[i] = make([]bool, width)
	}
	return new_u
}

// method Show prints the grid ; * for live, ' '  for dead cells
func (u Universe) Show() {
	for i := range u {
		for x := range u[i] {
			if u[i][x] == true {
				fmt.Printf("*")
			} else {
				fmt.Printf("-")
			}
		}
		fmt.Println()
	}
}

// Seed randomly initializes ~25% of cells
func Seed(u Universe) {
	rand.Seed(time.Now().UnixNano())
	for i := range u {
		for x := range u[i] {
			random := rand.Intn(4)
			if random == 1 {
				u[i][x] = true
			}
		}
	}
}

// Alive method returns whether a cell is alive or dead
// checks whether x, y are valid and wraps around
// see comment at end of program for why x and y are reversed
func (u Universe) Alive(y, x int) bool {
	if y < 0 {
		y += height
	}
	if y >= height {
		y -= height
	}
	if x < 0 {
		x += width
	}
	if x >= width {
		x -= width
	}
	return u[y][x]

}

func (u Universe) Neighbors(y, x int) int {
	var count int = 0
	// outer for loop jumps "rows" / y
	for i := y - 1; i <= y+1; i++ {
		// inner for loop checks "columns" / x
		for l := x - 1; l <= x+1; l++ {
			// don't count cell in question
			if i == y && x == l {
				continue
			}
			// otherwise check every neighbor and count
			if u.Alive(i, l) {
				count++
			}
		}

	}

	return count
}

// func Next computes the next state of a given cell
func (u Universe) Next(y, x int) bool {
	is_alive := u.Alive(y, x)
	n := u.Neighbors(y, x)
	if is_alive {
		if n < 2 || n > 3 {
			return false
		} else {
			return true
		}
	} else {
		if n == 3 {
			return true
		} else {
			return false
		}
	}
	// else, if on separate lines is more readable here imo
}

func clear_screen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// starts conway's game of life given a universe
func start(u Universe) {
	u_next := NewUniverse()
	clear_screen()
	u.Show()
	for n := 0; n < generations; n++ {
		for i := range u {
			for l := range u[i] {
				u_next[i][l] = u.Next(i, l)
			}
		}
		u = u_next
		u_next = NewUniverse()
		time.Sleep(200 * time.Millisecond)
		clear_screen()
		u.Show()
	}
}

func main() {

	u := NewUniverse()

	//Glider
	//u[2][4] = true
	//u[3][5] = true
	//u[4][3] = true
	//u[4][4] = true
	//u[4][5] = true

	Seed(u)
	start(u)

}

/*
calling functions with x and y as parameters for height and width is unfortunate
with our slice like so: sl[height][width]
x ends up being height / rows
y ends up being width / columns
this is completely counter to as x and y appear on paper when having a 2d coordinate system
so I thought it would make more sense to have y, x as parameters for height and width
*/
