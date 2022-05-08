package main

import (
	"fmt"
)

func main() {

	go func_a()
	go func_b()
	go func_c()

	// as a trick: an empty select waits forever
	// might be helpful if you have goroutines running and want to run them indefinitely
	// otherwise end of main() will end all goroutines

	select {}
	// in the end select{} detects no running goroutines at which point it
	// assumes deadlock, so it will end in a fatal error here

}

func func_a() { fmt.Println("a") }
func func_b() { fmt.Println("b") }
func func_c() { fmt.Println("c") }
