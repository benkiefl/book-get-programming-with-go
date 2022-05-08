// let's assume we have a rover driving around mars
// and represent an engine with a goroutine updating a position
// this is how we will implement it:
// 		the image package provides a Point type (a struct with X, Y as integers)
//		using a typical worker layout (often a for loop using select)
//			(in other langs: event loop, in go: goroutines are used)
//		we can use time.After to awake every so often to update the position

package main

import (
	"fmt"
	"image"
	"time"
)

func main() {

	go worker()
	time.Sleep(5 * time.Second)

}

func worker() {
	location := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}
	next := time.After(500 * time.Millisecond)
	for {
		select {
		case <-next:
			location = location.Add(direction)
			fmt.Printf("current position is %v\n", location)
			next = time.After(500 * time.Millisecond)
		}
	}

}
