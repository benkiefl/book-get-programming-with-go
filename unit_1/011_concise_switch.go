package main

import "fmt"

func main() {

	var command string = "south"

	switch command {
	case "north": // case command == "north"
		fmt.Println("You head north")
	case "east":
		fmt.Println("You head east")
	case "west":
		fmt.Println("You head west")
	case "south":
		fmt.Println("You head south")
		fallthrough // next case will be executed
	case "todemonstrate_fallthrough":
		fmt.Println("Where it's warmer")
	default:
		fmt.Println("illegal input")

	}

	// C, Java, Javascript fall through by default
	// so break; statements are used
	// in Go fallthrough is "explicit"

}
