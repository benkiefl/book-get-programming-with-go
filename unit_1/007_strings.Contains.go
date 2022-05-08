package main

import (
	"fmt"
	"strings"
)

func main() {

	var something string = "outside the woods the trees are standing tall"

	fmt.Println(strings.Contains(something, "out"))
	fmt.Println(strings.Contains(something, "xyz"))

}
