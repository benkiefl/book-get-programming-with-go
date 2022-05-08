package main

import (
	"fmt"
	"time"
)

func main() {

	// channels are used to send values between goroutines
	// channels are created using make
	c := make(chan int)

	// the close function closes a channel
	close(c)

	// the go statement starts a new goroutine (running concurrently)
	go func(c chan int) {
		fmt.Println(<-c) // the <- operator is used to send/receive
		// reading from a closed channel: corresponding zero value
		// writing to a closed channel panics
	}(c)
	time.Sleep(1 * time.Second) // when hitting end of main goroutines will shut down

	// the comma ok syntax can be used to check whether a channel is closed
	val, ok := <-c
	fmt.Printf("%v %v\n", val, ok)

	// though range can be used to automatically read from a channel until it's closed

	c1 := make(chan string, 10)
	defer close(c1)
	words := []string{"hello", "darkness", "my", "old", "friend"}

	go func() {
		for _, word := range words {
			c1 <- word
		}
	}()

	go func() {
		for word := range c1 {
			fmt.Printf("%s ", word)
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println()

}
