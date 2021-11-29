package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("2020/06/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(first(input))
	fmt.Println(second(input))
}

func first(input []byte) int {
	count := 0
	for _, inputString := range strings.Split(string(input), "\n\n") {
		answers := make(map[rune]int)
		for _, char := range inputString {
			if char != '\n' {
				answers[char]++
			}
		}
		count += len(answers)
	}
	return count
}

func second(input []byte) int {
	count := 0
	for _, inputString := range strings.Split(string(input), "\n\n") {
		must := len(strings.Split(inputString, "\n"))
		answers := make(map[rune]int)
		for _, char := range inputString {
			if char != '\n' {
				answers[char]++
			}
		}
		for _, given := range answers {
			if given == must {
				count += 1
			}
		}
	}
	return count
}
