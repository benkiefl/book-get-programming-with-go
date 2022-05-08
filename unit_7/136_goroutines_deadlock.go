package main

func main() {

	// when one or more goroutines end up waiting for something that will never happen
	// it's called a deadlock

	c := make(chan int)
	<-c // trying to recieve from a channel, program execution locks up
}

// in large programs deadlocks can involve an intricate series of dependencies
// between goroutines
