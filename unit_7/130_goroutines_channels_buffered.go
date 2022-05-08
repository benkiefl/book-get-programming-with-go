package main

import "fmt"

func main() {

	// let's create a buffered channel with buffer capacity 1
	c := make(chan int, 1)

	go increment(c)

	// since c is now a buffered channel, when we send something to it, it will
	// not wait for the item to be received on the other end.
	c <- 99
	// integer 99 will be stored in the channel, execution will immediately continue

	i := <-c

	fmt.Println(i) // i will most likely be 99

}

func increment(c chan int) {

	r := <-c
	fmt.Println("never prints")
	r++
	c <- r

}

// the execution flow doesn't have to be that exact way. Sending to / receiving
// from the channel are never properly synched, besides that a lot can happen.
// So "never prints" line never printing, i being 99, though on my system that
// always being the case, it doesn't have to be

// the golang playground sometimes gave a panic: send on closed channel
// run https://play.golang.org/p/XLHZoALWllD multiple times
