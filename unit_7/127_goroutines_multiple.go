// goroutines appear to run at the same time
// behind the scenes though there is a lot of possible context switching
// it's not necessarily strictly parallel
// (runtime, OS, CPU, number of CPU threads, a lot of variants influence that)
// therefore: it's best to assume that order is not deterministic and in fact it isn't

package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 5; i++ {
		go func_a()
		go func_b()
		go func_c()
		go func_d()

		time.Sleep(5 * time.Second)
		fmt.Println("end of loop")
	}
}

func func_a() {
	time.Sleep(3 * time.Second)
	fmt.Println("func_a")
}

func func_b() {
	time.Sleep(3 * time.Second)
	fmt.Println("func_b")
}

func func_c() {
	time.Sleep(3 * time.Second)
	fmt.Println("func_c")
}

func func_d() {
	time.Sleep(3 * time.Second)
	fmt.Println("func_d")
}
