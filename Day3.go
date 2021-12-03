package main

import (
	"fmt"

	"github.com/MaxdenHartog/AoC2021/helpers"
)

func (function *Function) Day3_1() {
	data, length := helpers.ReadLines("input/03.txt")

	ones := make([]int, len(data[0]))

	gamma := ""

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
		} else {
			gamma += "0"
		}
	}

	fmt.Println(gamma)
}
