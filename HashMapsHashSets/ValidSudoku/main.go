package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	strMap := make([]map[byte]int, 9)
	colMap := make([]map[byte]int, 9)
	for i := 0; i < 9; i++ {
		strMap[i] = make(map[byte]int)
		colMap[i] = make(map[byte]int)
	}
	squares := make([][]map[byte]int, 3)
	for i := range squares {
		squares[i] = make([]map[byte]int, 3)
		for j := range squares[i] {
			squares[i][j] = make(map[byte]int)
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			val := board[i][j]
			if val != '.' {
				if _, ex := strMap[i][byte(val)]; ex {
					return false
				} else {
					strMap[i][byte(val)]++
				}

				if _, ex := colMap[j][byte(val)]; ex {
					return false
				} else {
					colMap[j][byte(val)]++
				}

				iSquare := i / 3
				jSquare := j / 3
				if _, ex := squares[iSquare][jSquare][byte(val)]; ex {
					return false
				} else {
					squares[iSquare][jSquare][byte(val)]++
				}
			} else {
				continue
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}))
}
