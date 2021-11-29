package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("2020/01/input.txt")
	if err != nil {
		panic(err)
	}

	var values []int
	for _, inputString := range strings.Split(string(input), "\n") {
		inputInt, err := strconv.Atoi(inputString)
		if err != nil {
			panic(err)
		}

		values = append(values, inputInt)
	}

	fmt.Println(first(values))
	fmt.Println(second(values))
}

func first(values []int) int {
	for _, a := range values {
		for _, b := range values {
			if a+b == 2020 {
				return a * b
			}
		}
	}
	return -1
}

func second(values []int) int {
	for _, a := range values {
		for _, b := range values {
			for _, c := range values {
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}
	return -1
}
