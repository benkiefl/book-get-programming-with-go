package main

import (
	"fmt"
	"time"
)

func main() {

	// select is like a switch for channels
	// each case holds a channel receive or send

	c := make(chan int)
	d := make(chan int)

	go func_c(c) // pass channel c to func_c
	go func_d(d) // pass channel d to func_d

	c <- 0                      // send 0 to channel c
	d <- 5                      // send 5 to channel d
	time.Sleep(1 * time.Second) // sleep or otherwise default case in select will run

	// select waits till one case is ready, then runs it
	// if no channel is ready, select will run default
	select {
	// channel c receives from func_c which sleeps for two seconds
	case c_val := <-c:
		fmt.Println(c_val)
	// channel d receives from func_d which doesn't sleep, channel will be ready quicker
	case d_val := <-d:
		fmt.Println(d_val)
	default:
		fmt.Println("in case no channels are ready")
	}
}


// func_c and func_d take a channel, recieve a value, increment it, and send
// it back to the channel

func func_c(c chan int) {
	r := <-c
	r++
	time.Sleep(2 * time.Second)
	c <- r
}
func func_d(d chan int) {
	r := <-d
	r++
	d <- r
}
