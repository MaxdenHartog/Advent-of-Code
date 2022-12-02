package main

import (
	"fmt"
	"strconv"

	"github.com/MaxdenHartog/AoC2021/go/helpers"
)

func (function *Function) Day2_1() {
	data, length := helpers.ReadLines("input/02.txt")

	instructions := make(map[string]int)

	for i := 0; i < length; i++ {
		instruction := string(data[i][0])
		distance, _ := strconv.Atoi(string(data[i][len(data[i])-1]))

		instructions[instruction] += distance
	}

	result := instructions["f"] * (instructions["d"] - instructions["u"])

	fmt.Println(result)
}

func (function *Function) Day2_2() {
	data, length := helpers.ReadLines("input/02.txt")

	instructions := make(map[string]int)

	for i := 0; i < length; i++ {
		instruction := string(data[i][0])
		distance, _ := strconv.Atoi(string(data[i][len(data[i])-1]))

		if instruction == "f" {
			instructions["depth"] += distance * (instructions["d"] - instructions["u"])
		}

		instructions[instruction] += distance
	}

	result := instructions["f"] * instructions["depth"]

	fmt.Println(result)
}
