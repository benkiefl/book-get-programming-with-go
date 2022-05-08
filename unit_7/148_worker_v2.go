package main

import (
	"fmt"
	"image"
	"math/rand"
	"time"
)

type Rover struct {
	c chan command
}

type command int

const (
	right = command(0)
	left  = command(1)
)

func NewRover() *Rover {
	r := &Rover{c: make(chan command)}
	go r.Drive()
	return r
}

func (r *Rover) Drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 400 * time.Millisecond
	next := time.After(updateInterval)
	for {
		select {
		case c := <-r.c:
			switch c {
			case right:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
				fmt.Printf("right turn! ")
			case left:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
				fmt.Printf("left turn! ")
			}
			fmt.Printf("new direction %v\n", direction)
		case <-next:
			pos = pos.Add(direction)
			fmt.Printf("moved to %v\n", pos)
			next = time.After(updateInterval)
		}
	}
}

func main() {

	r := NewRover()

	rand.Seed(time.Now().UnixNano())
	for {
		time.Sleep(2 * time.Second)
		r.c <- command(rand.Intn(2))
	}

}
