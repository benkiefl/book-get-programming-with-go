// exact same code as example before

package main

import "fmt"

func main() {

	c := make(chan int) // same as c := make(chan int, 0)
	// above creates a unbuffered channel, that means the channel has no
	// "cache", so it can't store/hold any values

	go increment(c)

	// since the channel is unbuffered (think of it as capacity 0), as soon as
	// we send something to the channel, it will need to wait until that
	// something is taken out (received)

	// another way to think about it: we write integer 99 to the channel, but
	// the channel has no capacity to store the value, the value needs to be
	// received, so it can be passed

	c <- 99

	// so that's why sending to / receiving from the channel are synch up points
	// we send the integer 99 to the channel, it will wait until it's received,
	// in this case until it's received by increment()

	i := <-c // the same here, it's a definite synch up point with c <- r from increment()

	fmt.Println(i) // so i here will always be 100

}

func increment(c chan int) {

	r := <-c
	r++
	c <- r

}
