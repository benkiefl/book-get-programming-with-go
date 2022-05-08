package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // create an unbuffered channel
	// unbuffered channel: receive and send will synch up
	for i := 0; i < 5; i++ {
		// the goroutines run in a non-deterministic order
		go sleepyGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c // the receive will synch up with the send from sleepyGopher()
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	// sleeps, main will wait on gopherID := <-c for a value sent
	fmt.Println("... ", id, " snore ...")
	c <- id
}
