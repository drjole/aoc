package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("2015/01/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(first(input))
	fmt.Println(second(input))
}

func first(input []byte) int {
	result := 0
	for _, character := range input {
		if character == '(' {
			result += 1
		} else if character == ')' {
			result -= 1
		} else {
			panic("this should never happen")
		}
	}
	return result
}

func second(input []byte) int {
	result := 0
	for position, character := range input {
		if character == '(' {
			result += 1
		} else if character == ')' {
			result -= 1
		} else {
			panic("this should never happen")
		}
		if result < 0 {
			return position + 1
		}
	}
	return -1
}
