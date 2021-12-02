package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile("2021/01/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputBytes), "\n")

	fmt.Println(first(lines))
	fmt.Println(second(lines))
}

func first(lines []string) int {
	previous, _ := strconv.Atoi(lines[0])
	result := 0
	for _, line := range lines[1:] {
		current, _ := strconv.Atoi(line)
		if current > previous {
			result++
		}
		previous = current
	}
	return result
}

func second(lines []string) int {
	numbers := make([]int, len(lines))
	for i, line := range lines {
		a, _ := strconv.Atoi(line)
		numbers[i] = a
	}

	result := 0
	for i := 3; i < len(numbers); i++ {
		if numbers[i]+numbers[i-1]+numbers[i-2] > numbers[i-1]+numbers[i-2]+numbers[i-3] {
			result++
		}
	}
	return result
}
