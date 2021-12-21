package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MaxdenHartog/AoC2021/helpers"
)

// =================================== Puzzle 1 ===================================
func (function *Function) Day4_1() {
	data, _ := helpers.ReadLines("input/example.txt")

	numbers, boards := parseBingo(data)
	result := 0
	for _, number := range numbers {

		for _, board := range boards {
			if j := helpers.ContainsInt(board, number); j != -1 {
				board[j] = -board[j]
				board[len(board)-1]++
			}

			// Don't check winning condition if board hasn't marked 5 numbers yet
			if board[len(board)-1] < 5 {
				continue
			}

			if result = CheckWinningCondition(board); result != 0 {
				fmt.Println(result)
				fmt.Println(number)
				result = result * number
				break
			}
		}
	}

	for _, board := range boards {
		fmt.Println(board)
	}

	fmt.Println(result)
}

// =================================== Helpers ===================================
func parseBingo(data []string) (numbers []int, boards [][]int) {
	numberString := strings.Split(data[0], ",")

	// Parse drawn numbers
	for _, v := range numberString {
		number, _ := strconv.Atoi(v)
		numbers = append(numbers, number)
	}

	// Split the boards into separate slices
	b := 0
	boards = append(boards, make([]int, 0))
	for i := 1; i+1 < len(data); i++ {

		// Append a counter (starting at 0) to the end of each board
		if len(boards[b]) == 25 {
			boards[b] = append(boards[b], 0)
			boards = append(boards, make([]int, 0))
			b++
		}

		line := strings.Fields(data[i+1])
		for _, v := range line {
			number, _ := strconv.Atoi(v)
			boards[b] = append(boards[b], number)
		}
	}
	boards[b] = append(boards[b], 0)

	return numbers, boards
}

func CheckWinningCondition(board []int) (result int) {
	for _, v := range winningConditions {
		for _, i := range v {
			if board[i] >= 0 {
				result = 0
				break
			}
			result += board[i]
		}
	}

	return result * -1
}

var winningConditions = [][]int{{0, 1, 2, 3, 4}, {5, 6, 7, 8, 9}, {10, 11, 12, 13, 14}, {15, 16, 17, 18, 19}, {20, 21, 22, 23, 24},
	{0, 5, 10, 15, 20}, {1, 6, 11, 16, 21}, {2, 7, 12, 17, 22}, {3, 8, 13, 18, 23}, {4, 9, 14, 19, 24}}
