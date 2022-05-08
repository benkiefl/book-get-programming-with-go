// code can be made concurrent by starting it in a goroutine
// goroutines use channels for communication
// goroutines are similar to threads, processes in other languages
// (but not quite the same)

package main

import (
	"fmt"
	"time"
)

func main() {
	go function()               // is executed, execution of main continues
	time.Sleep(4 * time.Second) // sleep here since when we hit } all goroutines are stopped
} // end of scope stops the goroutine

func function() {
	time.Sleep(3 * time.Second)            // 1 second less than we wait in main(), therefore
	fmt.Println("print within function()") // Println will rune before main() is finished
}
