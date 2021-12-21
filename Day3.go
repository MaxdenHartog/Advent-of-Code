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

	oxygen := Day3_2_ImplementationOxygen(0, data)
	c02 := Day3_2_ImplementationC02(0, data, length)

	oxygenValue, _ := strconv.Atoi(oxygen)
	co2Value, _ := strconv.Atoi(c02)

	result := oxygenValue * co2Value

	fmt.Println(result)
}

func Day3_2_ImplementationOxygen(column int, data []string) (oxygen string) {
	ones := 0
	zeroes := 0
	oneBits := make([]string, len(data))
	zeroBits := make([]string, len(data))
	fmt.Println(len(data))
	j := 0
	for i := 0; i < len(data); i++ {
		fmt.Println(i)
		if data[i][column] == '1' {
			oneBits[j] = data[i]
			ones++
			j++
		} else {
			zeroBits[j] = data[i]
			zeroes++
			j++
		}
	}

	if len(data) > 1 {
		if ones >= zeroes {
			fmt.Println(len(oneBits))
			Day3_2_ImplementationOxygen(column+1, oneBits)
		} else {
			Day3_2_ImplementationOxygen(column+1, zeroBits)
		}
	}

	return oxygen
}

func Day3_2_ImplementationC02(column int, data []string, length int) (c02 string) {
	ones := 0
	zeroes := 0
	oneBits := make([]string, length)
	zeroBits := make([]string, length)

	for i := 0; i < len(data); i++ {
		if data[i][column] == '1' {
			oneBits[column] = data[i]
			ones++
		} else {
			zeroBits[column] = data[i]
			zeroes++
		}
	}

	if len(data) > 1 {
		if zeroes <= ones {
			Day3_2_ImplementationC02(column+1, zeroBits, length)
		} else {
			Day3_2_ImplementationC02(column+1, oneBits, length)
		}
	}

	return c02
}
