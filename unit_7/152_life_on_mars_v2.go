package main

import (
	"fmt"
	"image"
	"math/rand"
	"sync"
	"time"
)

const (
	rows = 20
	cols = 80
)

const (
	UPDATEROVER  = 50 * time.Millisecond
	TICKSPEEDECC = 100 * time.Millisecond
)

type Field struct {
	mu   sync.Mutex
	name string
	Grid [rows][cols]Cell
}

type Cell struct {
	life_chance int
	is_occupied bool
}

// NewField is a simple constructor function ; by using this constructor
// function for all celestial bodies we give each an equal chance of setting off
// life detectors, but I didn't want to over do it
func NewField(name string) *Field {
	f := Field{name: name}
	for i := range f.Grid {
		for l := range f.Grid[i] {
			rval := rand.Intn(100) + 1
			f.Grid[i][l].life_chance = rval
		}
	}
	return &f
}

type Rover struct {
	name   string
	chan_c chan command // command channel
	chan_m chan string  // message channel
	// buffer   []string
	which_cb string // which celestial body the rover is on
}

type command int

const (
	// stop   = command(0)
	// start  = command(1)
	right  = command(2)
	left   = command(3)
	status = command(4) // status of current rover, includes position
	send   = command(5) // send messages
	ping   = command(6) // simply sends a signal back
)

// NewRover is a simple constructor function for rovers, starts method Drive()
// for each rover
func NewRover(name string, f *Field) *Rover {
	r := Rover{
		name:     name,
		chan_c:   make(chan command),
		chan_m:   make(chan string, 30),
		which_cb: f.name,
	}
	go r.Drive(f)
	return &r
}

// Drive() method, invoked as a goroutine in constructorfunction NewRover; makes
// use of a for loop + select to keep the rover running given a 2-dimensional
// Grid as field
func (r *Rover) Drive(f *Field) {
	buffer := make([]string, 30)
	pos := image.Point{Y: rand.Intn(rows), X: rand.Intn(cols)}
	direction := image.Point{Y: 0, X: 1}
	next := time.After(UPDATEROVER)
	f.mu.Lock()

	// assign a new position until an empty one is found
	for f.Grid[pos.Y][pos.X].is_occupied {
		pos = image.Point{Y: rand.Intn(rows), X: rand.Intn(cols)}
	}
	f.Grid[pos.Y][pos.X].is_occupied = true

	// check if initial position has a life expectancty rating of above 99
	if f.Grid[pos.Y][pos.X].life_chance > 99 {
		buffer = append(buffer,
			fmt.Sprintf("life sensors triggered on %s, position %v by rover %s", r.which_cb, pos, r.name))
	}
	f.mu.Unlock()

	// now begins the main logic keeping the rovers driving around
	for {
		select {
		// here if we get a command to the rover
		case c := <-r.chan_c:
			switch c {
			case right:
				direction = image.Point{
					Y: direction.X,
					X: -direction.Y,
				}
				buffer = append(buffer,
					fmt.Sprintf("rover %s on %s turned right", r.name, r.which_cb))
			case left:
				direction = image.Point{
					Y: -direction.X,
					X: direction.Y,
				}
				buffer = append(buffer,
					fmt.Sprintf("rover %s on %s turned left", r.name, r.which_cb))
			case status:
				buffer = append(buffer,
					fmt.Sprintf("rover %s on %s currently positioned at %v, direction %v", r.name, r.which_cb, pos, direction))
			case ping:
				buffer = append(buffer,
					fmt.Sprintf("rover %s on %s pinged successfully", r.name, r.which_cb))
			case send:
				for _, message := range buffer {
					r.chan_m <- message // the point is to simulate sending via channels, so I did
				}
				close(r.chan_m)
			}
		// here if the rover is supposed to move
		case <-next:
			nextpos := pos.Add(direction)
			direction_new := direction
			f.mu.Lock()
			// our nextposition can't leave the matrix and can't be taken up by another rover
			for nextpos.Y < 0 || nextpos.X < 0 || nextpos.Y >= rows || nextpos.X >= cols || f.Grid[nextpos.Y][nextpos.X].is_occupied {
				direction_new = random_direction()
				nextpos = pos.Add(direction_new)
			}
			direction = direction_new
			f.Grid[pos.Y][pos.X].is_occupied = false
			pos = nextpos
			f.Grid[pos.Y][pos.X].is_occupied = true
			// life check
			if f.Grid[pos.Y][pos.X].life_chance > 99 {
				buffer = append(buffer,
					fmt.Sprintf("life sensors triggered on %s, position %v by rover %s", r.which_cb, pos, r.name))
			}
			f.mu.Unlock()
			next = time.After(UPDATEROVER)
		}
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	// celestial body 1 & 2
	cb1 := NewField("Mars")
	cb2 := NewField("Titan")

	// various rovers on Mars
	r1 := NewRover("Archimedes", cb1)
	r2 := NewRover("Newton", cb1)
	r3 := NewRover("Maxwell", cb1)
	// various rovers on Titan
	r4 := NewRover("Dirac", cb2)
	r5 := NewRover("Rutherford", cb2)

	// Print an Overview
	fmt.Printf("Our rovers are on %s and %s\n", cb1.name, cb2.name)
	fmt.Printf("Rover\t\tLocation\n")
	fmt.Printf("-----\t\t--------\n")
	fmt.Printf("%s\t%s\n", r1.name, r1.which_cb)
	fmt.Printf("%s\t\t%s\n", r2.name, r2.which_cb)
	fmt.Printf("%s\t\t%s\n", r3.name, r3.which_cb)
	fmt.Printf("%s\t\t%s\n", r4.name, r4.which_cb)
	fmt.Printf("%s\t%s\n", r5.name, r5.which_cb)
	fmt.Println()

	// Start Earth Control Station (function ECC())
	fmt.Printf("Earth Control Center starting... ")
	go ECS(r1, r2, r3, r4, r5, cb1, cb2)

	// wait time before program terminates
	time.Sleep(20 * time.Second)
}

// ECS is our Earth Control Station; sends commands to rovers within a
// pre-defined window; constant TICKSPEEDECC is one hour; times rovers are
// reachable are hard coded, in case of Titan it only happens every second day
func ECS(r1, r2, r3, r4, r5 *Rover, cb1, cb2 *Field) {
	// commented out Printf statements were used for debugging
	fmt.Printf("ECS started successfully\n")
	for hour, day := 0, 1; ; hour++ {
		// mars window
		if hour >= 5 && hour < 6 {
			if day == 1 {
				//fmt.Printf("day %2d hour %2d: pinging rover 1, 2, 3\n", day, hour)
				r1.chan_c <- ping
				r2.chan_c <- ping
				r3.chan_c <- ping
			}
			if day == 3 {
				//fmt.Printf("day %2d hour %2d: left turn rover 1 and 3, right turn rover 2\n", day, hour)
				r1.chan_c <- left
				r2.chan_c <- right
				r3.chan_c <- left
			}
			if day == 5 {
				//fmt.Printf("day %2d hour %2d: status check on rover 1\n", day, hour)
				r1.chan_c <- status
			}
			if day == 7 {
				//fmt.Printf("day %2d hour %2d: telling rover 1, 2, 3 to start sending\n", day, hour)
				r1.chan_c <- send
				r2.chan_c <- send
				r3.chan_c <- send
			}
		}
		// titan window
		if day%2 == 0 && hour >= 16 && hour < 17 {
			//fmt.Printf("day %2d hour %2d: pinging rover 4, 5 ; status check on rover 4, 5\n", day, hour)
			r4.chan_c <- ping
			r5.chan_c <- ping
			r4.chan_c <- status
			r5.chan_c <- status
		}
		if day == 7 && hour >= 16 && hour < 17 {
			//fmt.Printf("day %2d hour %2d: telling rover 4, 5 to start sending\n", day, hour)
			r4.chan_c <- send
			r5.chan_c <- send
		}
		// if 24 hours reached, reset hours, increment day
		if hour/24 == 1 {
			// print current rover positions once a day; you can print those at
			// end of every ECS cycle if you want to see hit detection working
			fmt.Println("Current Rover positions on Mars")
			print_Grid(cb1)
			fmt.Println("Current Rover positions on Titan")
			print_Grid(cb2)
			hour = 0
			day++
		}
		if day > 7 {
			break
		}
		time.Sleep(TICKSPEEDECC)
	}

	// in theory I could only retrieve in the given time windows as well, but I
	// mean, when is it just too much, it's just an exercise

	Retrieve(r1)
	Retrieve(r2)
	Retrieve(r3)
	Retrieve(r4)
	Retrieve(r5)
}

// random_direction returns a random direction (up, left, right, down)
func random_direction() image.Point {
	up := image.Point{Y: 1, X: 0}
	right := image.Point{Y: 0, X: 1}
	left := image.Point{Y: 0, X: -1}
	down := image.Point{Y: -1, X: 0}

	rval := rand.Intn(4)
	switch rval {
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

// print_Grid prints a given field, X for true, - for empty
func print_Grid(field *Field) {
	for i := range field.Grid {
		for l := range field.Grid[i] {
			if field.Grid[i][l].is_occupied {
				fmt.Printf("X")
			} else {
				fmt.Printf("-")
			}
		}
		fmt.Println()
	}

}

// Retrieve sends a send signal to a given rover, ranges over the message
// channel of the rover to print out all the messages
func Retrieve(r *Rover) {
	fmt.Printf("All messages from rover %s\n", r.name)
	for message := range r.chan_m {
		if !(message == "") {
			fmt.Printf("--- %v\n", message)
		}
	}
}
