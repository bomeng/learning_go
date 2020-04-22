/*
36. Valid Sudoku

Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.
The Sudoku board could be partially filled, where empty cells are filled with the character '.'.
*/
package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	var arr = [9]byte{}
	// check row
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			arr[col] = board[row][col]
		}
		if !isValid(arr) {
			return false
		}
	}

	// check column
	for col := 0; col < 9; col++ {
		for row := 0; row < 9; row++ {
			arr[row] = board[row][col]
		}
		if !isValid(arr) {
			return false
		}
	}

	// check square
	for row := 0; row < 9; row = row + 3 {
		for col := 0; col < 9; col = col + 3 {
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					arr[x*3+y] = board[row+x][col+y]
				}
			}
			if !isValid(arr) {
				return false
			}
		}
	}

	return true
}

func isValid(board [9]byte) bool {
	set := map[byte]bool{}
	for _, v := range board {
		if v != '.' {
			_, exist := set[v]
			if exist {
				return false
			} else {
				set[v] = true
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println(isValidSudoku(board))
}
