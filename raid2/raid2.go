package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	
	param:=strings.Join(os.Args[1:],"")

	if iscorrectparam(param){
		board := parseInput(param)
		if backtrack(&board) {
			if isValid(&board) && isBoardValid(&board){
				printBoard(board)
			}else {
				fmt.Printf("Error")
			}
		} else {
			fmt.Printf("Error")
		}
	}else{
		fmt.Println("Error")
	}

	
}

func backtrack(board *[9][9]int) bool {
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = candidate
					if isBoardValid(board) {
						if backtrack(board) {
							return true
						}
						board[i][j] = 0
					} else {
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func hasEmptyCell(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func isBoardValid(board *[9][9]int) bool {
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func isValid(board *[9][9]int) bool {
	for row := 0; row < 9; row++ {
		counter := 0
		for col := 0; col < 9; col++ {
			counter+=board[row][col]
		}
		if 45!= counter {
			return false
		}
	}

	for row := 0; row < 9; row++ {
		counter := 0
		for col := 0; col < 9; col++ {
			counter+=board[col][row]
		}
		if 45!= counter {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := 0
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter+=board[row][col]
				}
			}
			if 45!= counter {
				return false
			}
		}
	}

	return true
}

func parseInput(input string) [9][9]int {
	board := [9][9]int{}
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanRunes)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			scanner.Scan()
			i1, _ := strconv.Atoi(scanner.Text())
			board[row][col] = i1
		}
	}
	return board
}

func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func printBoard(board [9][9]int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Printf("%d", board[row][col])
			if col + 1 < 9 {
				fmt.Print(" ")
			} 
		}
		fmt.Println()
	}
}

func iscorrectparam(param string) bool{
	board := parseInput(param)
	
	if len(param)!=81{
		return false
	}
	
	for _,val:= range param{
		if !((48 <= val && val <= 57)||val==46){
			return false
		}
	}
	
	return isBoardValid(&board)
}
