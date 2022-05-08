package main

import (
	"fmt"
	"image"
	"time"
)

type Rover struct {
	mov chan command
}

type command int

const (
	stop  = command(0)
	start = command(1)
	right = command(2)
	left  = command(3)
	step  = 250 * time.Millisecond
)

func NewRover() *Rover {
	r := &Rover{mov: make(chan command)}
	go r.Drive()
	return r
}

func (r *Rover) Drive() {
	pos := image.Point{X: 0, Y: 0}
	direct := image.Point{X: 1, Y: 0}
	direct_save := image.Point{}
	next := time.After(step)
	for {
		select {
		case c := <-r.mov:
			switch c {
			case stop:
				fmt.Printf("Rover stopped, ")
				direct_save = direct
				direct = image.Point{0, 0}
			case start:
				fmt.Printf("Rover started, ")
				direct = direct_save
			case left:
				fmt.Printf("Turning left, ")
				direct = image.Point{
					X: direct.Y,
					Y: -direct.X,
				}
			case right:
				fmt.Printf("Turning right, ")
				direct = image.Point{
					X: -direct.Y,
					Y: direct.X,
				}
			}
			fmt.Printf("new direction: %v\n", direct)
		case <-next:
			pos = pos.Add(direct)
			next = time.After(step)
			fmt.Printf("Position: %v\n", pos)
		}
	}
}

func main() {

	r := NewRover()

	time.Sleep(5 * step)
	r.mov <- right
	time.Sleep(5 * step)
	r.mov <- left
	time.Sleep(5 * step)
	r.mov <- stop
	time.Sleep(5 * step)
	r.mov <- start
	time.Sleep(5 * step)
}
