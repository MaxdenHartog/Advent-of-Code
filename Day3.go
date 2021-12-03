package main

import (
	"fmt"
	"strconv"

	"github.com/MaxdenHartog/AoC2021/helpers"
)

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

	// Couldn't figure out how to flip the bits so I resorted to this :(
	for i := 0; i < len(ones); i++ {
		if ones[i] > length-ones[i] {
			gamma += "1"
			epsilon += "0"
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
