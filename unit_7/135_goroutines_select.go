package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	c := make(chan int)

	go func(c chan int) {
		// we wait between 0 and 1999 miliseconds before sending to the channel
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		c <- 5
	}(c)

	timeout := time.After(1 * time.Second)

	select {
	case r := <-c:
		fmt.Println("successfully received from channel:", r)
	case <-timeout:
		fmt.Println("ran out of patience")
	}

}
