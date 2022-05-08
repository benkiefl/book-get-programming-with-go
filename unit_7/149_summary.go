// summary
//      never access state from more than one goroutine at a time unless documented
//		use mutexes to guard
//		select is like a switch on channels
//		long-lived goroutines can be written using select
//		hide worker details behind methods

package main
