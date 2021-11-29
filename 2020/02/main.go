package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("2020/02/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(first(input))
	fmt.Println(second(input))
}

func first(input []byte) int {
	r := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
	invalid := 0
	for _, inputString := range strings.Split(string(input), "\n") {
		match := r.FindStringSubmatch(inputString)
		minString, maxString, letter, password := match[1], match[2], match[3], match[4]
		min, err := strconv.Atoi(minString)
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(maxString)
		if err != nil {
			panic(err)
		}
		occurrences := strings.Count(password, letter)
		if min <= occurrences && occurrences <= max {
			invalid++
		}
	}
	return invalid
}

func second(input []byte) int {
	r := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
	valid := 0
	for _, inputString := range strings.Split(string(input), "\n") {
		match := r.FindStringSubmatch(inputString)
		minString, maxString, letter, password := match[1], match[2], match[3], match[4]
		min, err := strconv.Atoi(minString)
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(maxString)
		if err != nil {
			panic(err)
		}
		if (string(password[min-1]) == letter) != (string(password[max-1]) == letter) {
			valid++
		}
	}
	return valid
}
