package main

import "sync"

func main() {

	var mu sync.Mutex // unlocked by default

	// mu.Unlock()
	// fatal error: sync: unlock of unlocked mutex

	// mu.Lock()
	// mu.Lock()
	// fatal error: all goroutines are asleep - deadlock!
	// reason: we're waiting for the Lock to be given up while that never happens, deadlock

	// the more you do while the lock is held, the more you have to consider
	//		the lock might be held for a considerable amount of time
	//		we might easily deadlock
	// therefore
	//		keep code within mutex simple
	//		only have one mutex for a given piece of shared state

}
