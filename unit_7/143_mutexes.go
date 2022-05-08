// in Go there are shared values
// these can be used by multiple goroutines
// but this can lead to problems, if two or more simultaneously
// 		read from a shared value: works
//		read/write + write -"-: undefined behavior
// this is what leads to "race conditions" (routines racing to use the same value)

// mutex is a special type with a Lock and Unlock method (part of sync package)
// https://golang.org/pkg/sync/#Mutex
// mutex stands for mutual exclusion; typical workflow:
//      a goroutine wants to use the shared value
//      checks the mutex, if unlocked locks mutex, and uses shared value
//      no other goroutine uses shared value until
//			the mutex is unlocked by the goroutine that currently manipulates the value
// if one goroutine doesn't strictly follow protocol we might end up with a race condition
// that's why mutexes are almost always internal to a package

// if you are like me you're thinking "why not use a custom struct, with a bool
// for a check" or similar? but if you think it through the issue is quite
// complex given hundreds of possible goroutines and many shared values, how
// would your implementation handle goroutines waiting for a value to be
// available, and the value being handed off? how would you protect your
// repeated checks from being falsely optimized by the compiler (who might
// bundle them somehow)? et cetera

// be sure to check the source code https://golang.org/src/sync/mutex.go

package main

import "sync"

var mu sync.Mutex // the zero value is an unlocked Mutex

func main() {

	mu.Lock() // mutex locked
	defer mu.Unlock()
	// lock held until we return from function
}
