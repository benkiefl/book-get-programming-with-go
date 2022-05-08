package main

import "fmt"

func app(sl []string, num int) {

	for i := 0; i < num; i++ {
		cap_old := cap(sl)
		sl = append(sl, "a")
		cap_new := cap(sl)
		if cap_new > cap_old {
			fmt.Printf("cap increased from %10d to %10d\tmargin %v\n", cap_old, cap_new, cap_new-cap_old)
		}
	}
}

func main() {

	var sl []string

	app(sl, 100000)

}

// see https://golang.org/pkg/builtin/#cap for details
