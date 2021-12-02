package main

import (
	"fmt"

	"github.com/MaxdenHartog/AoC2021/helpers"
)

func (function *Function) Day1_1() {
	data, length := helpers.ReadLines("input/01.txt")
	parsedData := helpers.LinesAsInt(data)

	result := 0

	for i := 1; i < length; i++ {
		if parsedData[i] > parsedData[i-1] {
			result++
		}
	}

	fmt.Println(result)
}

func (function *Function) Day1_2() {
	data, length := helpers.ReadLines("input/01.txt")
	parsedData := helpers.LinesAsInt(data)

	result := 0
	window := parsedData[0] + parsedData[1] + parsedData[2]

	for i := 1; i < length-2; i++ {
		nextWindow := window - parsedData[i-1] + parsedData[i+2]
		if nextWindow > window {
			result++
		}

		window = nextWindow
	}

	fmt.Println(result)
}
