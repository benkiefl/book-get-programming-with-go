// one problem of course is how to know how long goroutines take (without sleep)
// go provides channels for goroutines to communciate and be synched
// channels can be thought of as pipes and just as any Go types, they can be used as
// variables, be passed, stored in structures et cetera see
// https://www.youtube.com/watch?v=f6kdp27TYZs

package main

import "fmt"

func main() {

	c := make(chan int) // this channel can send/receive only integers
	// you receive values via <- (left arrow operator)

	go increment(c) // function will execute and main will continue

	c <- 99 // sends the integer 99 into the channel

	// the send operation will wait until something (in another goroutine) tries
	// to receive on the same channel

	i := <-c

	fmt.Println(i)

}

func increment(c chan int) {

	r := <-c // go will wait until something is sent via the channel
	r++
	c <- r

}
