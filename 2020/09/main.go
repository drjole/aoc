package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile("2020/09/input.txt")
	if err != nil {
		panic(err)
	}
	input := string(inputBytes)

	fmt.Println(first(input))
	fmt.Println(second(input))
}

func first(input string) int {
	var numbers []int

	for _, numberString := range strings.Split(input, "\n") {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			panic(err)
		}

		if len(numbers) == 25 {
			found := false
			for _, a := range numbers {
				for _, b := range numbers {
					if a+b == number && a != b {
						found = true
					}
					if found {
						break
					}
				}
				if found {
					break
				}
			}
			if !found {
				return number
			}
		}

		if len(numbers) < 25 {
			numbers = append(numbers, number)
		} else {
			numbers = append(numbers[1:], number)
		}
	}
	return -1
}

func second(input string) int {
	target := first(input)
	sum := 0

	var numbers []int
	for _, numberString := range strings.Split(input, "\n") {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}

	var slidingWindow []int
	found := false
	for _, number := range numbers {
		for sum > target {
			sum -= slidingWindow[0]
			slidingWindow = slidingWindow[1:]
			if sum == target {
				found = true
			}
			if found {
				break
			}
		}
		if found {
			break
		}

		sum += number
		slidingWindow = append(slidingWindow, number)
		if sum == target {
			break
		}
	}

	min, max := math.MaxInt, 0
	for _, n := range slidingWindow {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min + max
}
