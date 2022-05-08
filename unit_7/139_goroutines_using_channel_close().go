package main

import (
	"fmt"
	"strings"
)

func main() {

	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher2(c0, c1)
	printGopher(c1)

}

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	close(downstream)
}

// filterGopher1 uses close and the comma, ok syntax
func filterGopher1(upstream, downstream chan string) {
	for {
		item, ok := <-upstream
		if !ok {
			close(downstream)
			return
		}
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
}

// filterGopher2 simply ranges over the channel ; this automatically reads from
// the channel until the channel is closed and can be thought of as a shortcut
// to checking that manually (like in filterGopher1())
func filterGopher2(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}
