package main

import "fmt"

func display(board [8][8]rune) {
	// range returns index and element
	// with an array[][] range gives you
	// the index, and the element is access to the rows
	for _, row := range board {
		for _, column := range row {
			if column == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c ", column)
			}
		}
		fmt.Println()
	}
}

func main() {

	var board [8][8]rune

	board[0] = [8]rune{'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r'}
	board[7] = [8]rune{'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R'}

	for index, _ := range board[1] {
		board[1][index] = 'p'
		board[6][index] = 'p'
	}

	display(board)
}
