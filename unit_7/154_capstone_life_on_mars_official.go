package main

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	dayLength         = 24 * time.Second // mars day
	receiveTimePerDay = 2 * time.Second  // receive window
)

const (
	right command = 0
	left  command = 1
)

type command int

// Message as sent from Mars (via chan Message)
type Message struct {
	Pos       image.Point
	LifeSigns int    // score from 1 to 1000
	Rover     string // rover name, rover{0..4} in this case
} // has no methods

// MarsGrid represents a grid on mars
type MarsGrid struct {
	bounds image.Rectangle
	mu     sync.Mutex // guarded
	cells  [][]cell
} // has methods Occupy, cell, Size

type cell struct {
	groundData SensorData // integer
	occupier   *Occupier
} // has no methods

type SensorData struct {
	LifeSigns int
} // has no methods

// RoverDriver represents the driving rover
type RoverDriver struct {
	commandc chan command
	occupier *Occupier
	name     string
	radio    *Radio
} // has methods left, right, drive, checkForLife

// radio which can send Message to earth
type Radio struct {
	fromRover chan Message
} // has methods run, SendToEarth

// Occupier, occupied cell; used to keep track of rovers
// mainly used via pointers (eg in type cell and RoverDriver)
// program heavily utilizes the fact pointer zero value is nil
type Occupier struct {
	grid *MarsGrid
	pos  image.Point
} // * has methods MoveTo, Sense, Pos

// run buffers messages sent by a rover until they can be sent to earth; this
// one was quite confusing at first
func (r *Radio) run(toEarth chan []Message) {
	var buffered []Message
	for {
		toEarth1 := toEarth // toEarth1 will refer to the same channel
		if len(buffered) == 0 {
			toEarth1 = nil
		}
		select {
		// write to nil channel can't happen, so in that case it will just wait for a message from channel r.fromRover
		case toEarth1 <- buffered:
			buffered = nil
		case m := <-r.fromRover:
			buffered = append(buffered, m)
		}
		// so if there's a message from rover, append to buffer, buffer will be non-zero, and messaage will be sent toEarth1
		// if buffer is empty again, toEarth1 will become nil, and select will again indefinitely wait for a further message
	}
}

// SendToEarth sends a message to earth
func (r *Radio) SendToEarth(m Message) {
	r.fromRover <- m // question why this has to be a separate method
}

// Left turns the rover left (90° counterclockwise).
func (r *RoverDriver) Left() {
	r.commandc <- left
}

// Right turns the rover right (90° clockwise).
func (r *RoverDriver) Right() {
	r.commandc <- right
}

// drive is responsible for driving the rover. It
// is expected to be started in a goroutine.
func (r *RoverDriver) drive() {
	log.Printf("%s initial position %v", r.name, r.occupier.Pos())
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			}
			log.Printf("%s new direction %v", r.name, direction)
		case <-nextMove:
			nextMove = time.After(updateInterval)
			newPos := r.occupier.Pos().Add(direction)
			if r.occupier.MoveTo(newPos) {
				log.Printf("%s moved to %v", r.name, newPos)
				r.checkForLife()
				break
			}
			log.Printf("%s blocked trying to move from %v to %v", r.name, r.occupier.Pos(), newPos)
			// Pick one of the other directions randomly.
			// Next time round, we'll try to move in the new
			// direction.
			dir := rand.Intn(3) + 1
			for i := 0; i < dir; i++ {
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			}
			log.Printf("%s new random direction %v", r.name, direction)
		}
	}
}

// Occupy occupies a cell at the given point in the grid. It
// returns nil if the point is already occupied or the point is outside
// the grid. Otherwise it returns a value that can be used
// to move to different places on the grid.
func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	g.mu.Lock()
	defer g.mu.Unlock()
	cell := g.cell(p) // if not in Bounds, nil, if in Bounds, cell will be of type *cell
	// if cell out of bounds || cell occupier (pointer) already used (in other words: occupied)
	if cell == nil || cell.occupier != nil {
		return nil
	}
	cell.occupier = &Occupier{
		grid: g,
		pos:  p,
	}
	return cell.occupier
}

func (g *MarsGrid) cell(p image.Point) *cell {
	// see https://golang.org/pkg/image/#Point.In
	if !p.In(g.bounds) {
		return nil
	}
	return &g.cells[p.Y][p.X]
}

// Size returns a Point representing the size of the grid.
func (g *MarsGrid) Size() image.Point {
	return g.bounds.Max
}

// MoveTo moves the occupier to a different cell in the grid.
// It reports whether the move was successful
// It might fail because it was trying to move outside
// the grid or because the cell it's trying to move into
// is occupied. If it fails, the occupier remains in the same place.
func (o *Occupier) MoveTo(p image.Point) bool {
	o.grid.mu.Lock()
	defer o.grid.mu.Unlock()
	newCell := o.grid.cell(p)
	if newCell == nil || newCell.occupier != nil {
		return false
	}
	o.grid.cell(o.pos).occupier = nil
	newCell.occupier = o
	o.pos = p
	return true
}

// Sense returns sensory data from the current cell.
func (o *Occupier) Sense() SensorData {
	o.grid.mu.Lock()
	defer o.grid.mu.Unlock()
	return o.grid.cell(o.pos).groundData
}

// Pos returns the current grid position of the occupier.
func (o *Occupier) Pos() image.Point {
	return o.pos
}

func main() {
	marsToEarth := make(chan []Message)
	go earthReceiver(marsToEarth)
	gridSize := image.Point{X: 20, Y: 10}
	grid := NewMarsGrid(gridSize)
	// this is very neat, create rover as a slice with pointers to various RoverDriver
	rover := make([]*RoverDriver, 5)
	for i := range rover {
		// range over rover, and startDriver for each, utilizing fmt.Sprint to
		// pass the index on as part of the name
		rover[i] = startDriver(fmt.Sprint("rover", i), grid, marsToEarth)
	}
	time.Sleep(10 * time.Second) // will be total run time
}

// earthReceiver receives from Mars on given channel, but only during receive
// time; function invoked in main() as goroutine
func earthReceiver(msgc chan []Message) {
	for {
		time.Sleep(dayLength - receiveTimePerDay)
		receiveMarsMessages(msgc)
	}
}

// receiveMarsMessages invoked by earthReceiver only in receive time frame
func receiveMarsMessages(msgc chan []Message) {
	finished := time.After(receiveTimePerDay)
	for {
		select {
		case <-finished:
			return
		case ms := <-msgc:
			// range over channel
			for _, m := range ms {
				log.Printf("earth received report of life sign level %d from %s at %v", m.LifeSigns, m.Rover, m.Pos)
			}
		}
	}
}

func startDriver(name string, grid *MarsGrid, marsToEarth chan []Message) *RoverDriver {
	var o *Occupier
	// zero value for a pointer is nil
	for o == nil {
		startPoint := image.Point{X: rand.Intn(grid.Size().X), Y: rand.Intn(grid.Size().Y)}
		o = grid.Occupy(startPoint) // Occupy the Grid with the rover at startpoint
	}
	return NewRoverDriver(name, o, marsToEarth)
}

// NewRadio returns a new Radio instance that sends
// messages on the toEarth channel.
func NewRadio(toEarth chan []Message) *Radio {
	r := &Radio{
		fromRover: make(chan Message),
	}
	go r.run(toEarth)
	return r
}

// NewRoverDriver initializes a RoverDriver, starts the drive() method as a goroutine
func NewRoverDriver(name string, occupier *Occupier, marsToEarth chan []Message) *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
		occupier: occupier,
		name:     name,
		radio:    NewRadio(marsToEarth),
	}
	go r.drive()
	return r
}

func (r *RoverDriver) checkForLife() {
	// Successfully moved to new position.
	sensorData := r.occupier.Sense()
	if sensorData.LifeSigns < 900 {
		return
	}
	r.radio.SendToEarth(Message{
		Pos:       r.occupier.Pos(),
		LifeSigns: sensorData.LifeSigns,
		Rover:     r.name,
	})
}

// NewMarsGrid returns a new MarsGrid of the given size.
func NewMarsGrid(size image.Point) *MarsGrid {
	grid := &MarsGrid{
		bounds: image.Rectangle{Max: size},
		cells:  make([][]cell, size.Y),
	}
	for y := range grid.cells {
		grid.cells[y] = make([]cell, size.X)
		for x := range grid.cells[y] {
			cell := &grid.cells[y][x]
			cell.groundData.LifeSigns = rand.Intn(1000)
		}
	}
	return grid
}
