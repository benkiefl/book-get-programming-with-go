package main

import (
	"fmt"
)

func main() {

	var b bool // default value: false
	var s string = "2"

	switch s {
	case "true", "yes", "0":
		b = true
	case "false", "no", "1":
		b = false
	default:
		fmt.Println("illegal input")
	}

	fmt.Println(b)

}
