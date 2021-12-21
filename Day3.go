package main

import (
	"fmt"
	"strconv"

	"github.com/MaxdenHartog/AoC2021/helpers"
)

// =================================== Puzzle 1 ===================================

func (function *Function) Day3_1() {
	data, length := helpers.ReadLines("input/03.txt")

	ones := make([]int, len(data[0]))

	gamma := ""
	epsilon := ""

	for i := 0; i < length; i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == '1' {
				ones[j]++
			}
		}

	}

	for i := 0; i < len(ones); i++ {
		if ones[i] > length-ones[i] {
			gamma += "1"
			epsilon += "0" // Couldn't figure out how to flip the bits so I resorted to this :(
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaValue, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonValue, _ := strconv.ParseInt(epsilon, 2, 64)
	result := gammaValue * epsilonValue

	fmt.Println(result)
}

// =================================== Puzzle 2 ===================================
func (function *Function) Day3_2() {
	data, length := helpers.ReadLines("input/03.txt")

	oxygen := Day3_2_ImplementationOxygen(0, data, length)
	co2 := Day3_2_ImplementationC02(0, data, length)

	result := oxygen * co2

	fmt.Println(result)
}

func Day3_2_ImplementationOxygen(column int, data []string, length int) (oxygen int64) {
	if length == 1 {
		oxygen, _ = strconv.ParseInt(data[0], 2, 64)
		return oxygen
	}

	ones := 0
	zeroes := 0
	oneBits := make([]string, length)
	zeroBits := make([]string, length)

	for i := 0; i < length; i++ {
		if data[i][column] == '1' {
			oneBits[ones] = data[i]
			ones++
		} else {
			zeroBits[zeroes] = data[i]
			zeroes++
		}
	}

	if ones >= zeroes {
		return Day3_2_ImplementationOxygen(column+1, oneBits, ones)
	} else {
		return Day3_2_ImplementationOxygen(column+1, zeroBits, zeroes)
	}
}

func Day3_2_ImplementationC02(column int, data []string, length int) (co2 int64) {
	if length == 1 {
		co2, _ = strconv.ParseInt(data[0], 2, 64)
		return co2
	}

	ones := 0
	zeroes := 0
	oneBits := make([]string, length)
	zeroBits := make([]string, length)

	for i := 0; i < length; i++ {
		if data[i][column] == '1' {
			oneBits[ones] = data[i]
			ones++
		} else {
			zeroBits[zeroes] = data[i]
			zeroes++
		}
	}

	if zeroes <= ones {
		return Day3_2_ImplementationC02(column+1, zeroBits, zeroes)
	} else {
		return Day3_2_ImplementationC02(column+1, oneBits, ones)
	}
}
