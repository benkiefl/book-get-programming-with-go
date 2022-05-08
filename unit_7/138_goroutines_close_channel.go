// you always can use a custom struct as sentinel value (eg with a boolean field)
// but another way is: you can close a channel

package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	c1 := make(chan float64)
	c2 := make(chan byte)

	close(c)
	close(c1)
	close(c2)

	// you can´t write to a closed channel
	// c <- 5 // panic: send on closed channel

	// if you read from a closed channel it will return the zero value for the type

	fmt.Println(<-c)
	fmt.Printf("%f %T\n", <-c1, <-c1)
	fmt.Printf("%b %T\n", <-c2, <-c2)

	// you can test whether a channel is closed with the comma ok idiom
	// (which is a two-valued assignement statement)
	_, test := <-c
	fmt.Printf("%v %T\n", test, test)

}
