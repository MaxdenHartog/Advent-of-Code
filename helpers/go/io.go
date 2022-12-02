package helpers

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func ReadFile(path string) (data []byte) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Printf("Failed reading file with path: %s", path)
	}

	return data
}

func ReadLines(path string) (data []string, lines int) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Printf("Failed reading file with path: %s", path)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	lines = len(data)

	return data, lines
}

func LinesAsInt(data []string) (parsed []int) {
	for i := 0; i < len(data); i++ {
		p, _ := strconv.Atoi(data[i])
		parsed = append(parsed, p)
	}

	return parsed
}
