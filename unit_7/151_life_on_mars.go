// this is 8 rovers randomly driving around and the field being printed out; it
// does not implement most of the features demanded by the exercise. consider it
// me playing around, though: hit detection works

package main

import (
	"fmt"
	"image"
	"math/rand"
	"os"
	"os/exec"
	"sync"
	"time"
)

const row = 40
const col = 80

type MarsGrid struct {
	mu     sync.Mutex
	Matrix [row][col]bool // true if rover present, false if not
}

// struct rover is empty; I realized that continously printing out the Field
// would complicate things with regards to channels, so I purged them from this
// version of the program
type Rover struct {
}

type command int

// NewRover() is our Constructor Function for a rover. Starts a goroutine
// Drive() (method of Rover)
func NewRover(m *MarsGrid) *Rover {
	r := &Rover{}
	go r.Drive(m)
	return r
}

// Drive() gives our rover a position, and makes sure it's driven around. Avoids
// hitting other rovers or leaving the matrix, most likely very UGLY code, even
// uses goto twice
func (r *Rover) Drive(m *MarsGrid) {
	// let the initial rover position be randomly on the Matrix
	// Y: row, X: col
	pos := image.Point{Y: rand.Intn(row), X: rand.Intn(col)}
	// remember: read + read/write at same time: undefined behavior
	// so we have to Lock() our mutex here
	m.mu.Lock()
	// if a position is already taken (True)
	for m.Matrix[pos.Y][pos.X] {
		// choose a new one until one is found
		pos = image.Point{Y: rand.Intn(row), X: rand.Intn(col)}
	}
	// initialize the position as true
	m.Matrix[pos.Y][pos.X] = true
	m.mu.Unlock()
	// this resembles past files, see 148_worker_v2.go
	UpdateInterval := 500 * time.Millisecond
	next := time.After(UpdateInterval)
	direc := random_direction()
	for {
		select {
		case <-next:
			// calculate next position given current direction
			nextpos := pos.Add(direc)
			m.mu.Lock()
			for {
			REPEAT:
				// check if our nextposition would leave our matrix
				if nextpos.Y >= row || nextpos.X >= col || nextpos.Y < 0 || nextpos.X < 0 {
					// if so, use a new random direction()
					direc = random_direction()
					// and calculate new position based on new direction
					nextpos = pos.Add(direc)
					// REPEAT the check
					goto REPEAT
				}
				// if nextposition is inside matrix, there might be another rover
				for m.Matrix[nextpos.Y][nextpos.X] {
					// if there is one, new random direction
					direc = random_direction()
					nextpos = pos.Add(direc)
					// REPEAT the checks (are we still inside the Matrix given our evasive maneuver)
					goto REPEAT
				}
				// if all is fine, leave the for loop
				break
			}
			// set current position to false
			m.Matrix[pos.Y][pos.X] = false
			pos = nextpos
			// set new position to true
			m.Matrix[pos.Y][pos.X] = true
			// print the Grid
			PrintGrid(m)
			m.mu.Unlock() // don't forget to unlock the mutex
			// reset "timer"
			next = time.After(UpdateInterval)
		}
	}
}

// PrintGrid prints our MarsGrid, "X" for true, "-" for false
func PrintGrid(m *MarsGrid) {
	clear_screen()
	for i := range m.Matrix {
		for l := range m.Matrix[i] {
			if m.Matrix[i][l] {
				fmt.Printf("X")
			} else {
				fmt.Printf("-")
			}
		}
		fmt.Println()
	}
}

// clear_screen() clears the screen
func clear_screen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func main() {

	Field := MarsGrid{}

	rover0 := NewRover(&Field)
	rover1 := NewRover(&Field)
	rover2 := NewRover(&Field)
	rover3 := NewRover(&Field)
	rover4 := NewRover(&Field)
	rover5 := NewRover(&Field)
	rover6 := NewRover(&Field)
	rover7 := NewRover(&Field)

	// if we don't do anything with the rovers, go will complain
	// unused variables
	fmt.Println(rover0, rover1, rover2, rover3,
		rover4, rover5, rover6, rover7)

	// let our rovers drive around for a given time
	time.Sleep(1 * time.Minute)

}

// random_direction returns a random direction
func random_direction() image.Point {
	up := image.Point{X: 0, Y: 1}
	right := image.Point{X: 1, Y: 0}
	left := image.Point{X: -1, Y: 0}
	down := image.Point{X: 0, Y: -1}
	switch rand.Intn(4) {
	case 0:
		return up
	case 1:
		return right
	case 2:
		return left
	case 3:
		return down
	default:
		return right
	}
}
