package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("2020/03/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(first(input))
	fmt.Println(second(input))
}

func first(input []byte) int {
	width := 0
	column := 0
	trees := 0
	for i, inputString := range strings.Split(string(input), "\n") {
		if i == 0 {
			width = len(inputString)
		} else {
			column = (column + 3) % width
			if string(inputString[column]) == "#" {
				trees++
			}
		}
	}
	return trees
}

func second(input []byte) int {
	prod := 1
	inputStrings := strings.Split(string(input), "\n")
	for _, slope := range []struct {
		right int
		down  int
	}{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	} {
		width := 0
		column := 0
		trees := 0
		for row := 0; row < len(inputStrings); row += slope.down {
			if row == 0 {
				width = len(inputStrings[row])
			} else {
				column = (column + slope.right) % width
				if string(inputStrings[row][column]) == "#" {
					trees++
				}
			}
		}
		prod *= trees
	}
	return prod
}
