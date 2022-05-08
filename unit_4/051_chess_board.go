// possibly complicated, error prone, over engineered, stupid way
// to create and display a chess_board [8][8] rune, I considered it an exercise

package main

import "fmt"

func cap_array(line [8]rune) [8]rune {
	var cap_line [8]rune
	var temp_rune rune
	for i := 0; i <= 7; i++ {
		temp_rune = line[i]
		// see https://en.wikipedia.org/wiki/List_of_Unicode_characters#Basic_Latin
		// on why it's -32
		cap_line[i] = rune(temp_rune - 32)
	}
	return cap_line
}

func main() {

	var chess_board [8][8]rune

	chess_line_1 := [8]rune{'r', 'n', 'b', 'k', 'q', 'b', 'n', 'r'}
	chess_line_2 := [8]rune{'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p'}
	chess_empty_line := [8]rune{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}

	chess_board[0] = chess_line_1
	chess_board[1] = chess_line_2

	for i := 2; i <= 5; i++ {
		chess_board[i] = chess_empty_line
	}

	chess_board[6] = cap_array(chess_line_2)
	chess_board[7] = cap_array(chess_line_1)

	for j := 0; j <= 7; j++ {
		for i := 0; i <= 7; i++ {
			fmt.Printf("%c ", chess_board[j][i])
		}
		fmt.Println()
	}
}
