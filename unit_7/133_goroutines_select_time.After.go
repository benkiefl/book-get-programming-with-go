package main

import (
	"fmt"
	"time"
)

func main() {

	timeout := time.After(2 * time.Second)
	c := make(chan int)

	// you might want to time out if no channel sends
	// you can do this via above timeout variable

	// channel c is never sent to

	select {
	case x := <-c: // channel c won't receive anything
		fmt.Println(x)
	case <-timeout:
		fmt.Println("we ran out of patience")
	}
}

// channels can be nil, in fact, nil is their default zero value
// as seen above: trying to use a nil channel won't panic
// instead the operation will block forever
// exception: close will panic a nil channel

// blocking does not take up any resources (except some memory while idling)
// (compared to an empty for loop for example which has your CPU racing)
