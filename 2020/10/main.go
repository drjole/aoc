package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile("2020/10/input.txt")
	if err != nil {
		panic(err)
	}
	input := string(inputBytes)

	fmt.Println(first(input))
	fmt.Println(second(input))
}

func first(input string) int {
	numbers := []int{0}
	for _, numberString := range strings.Split(input, "\n") {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	numbers = append(numbers, numbers[len(numbers)-1]+3)

	diffOne, diffThree := 0, 0
	for index := range numbers[:len(numbers)-1] {
		difference := numbers[index+1] - numbers[index]
		if difference == 1 {
			diffOne++
		}
		if difference == 3 {
			diffThree++
		}
	}
	return diffOne * diffThree
}

func second(input string) int {
	numbers := []int{0}
	for _, numberString := range strings.Split(input, "\n") {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	numbers = append(numbers, numbers[len(numbers)-1]+3)

	return pathsToEnd(numbers, 0)
}

var cache map[int]int

func init() {
	cache = make(map[int]int)
}

func pathsToEnd(numbers []int, start int) int {
	if value, ok := cache[start]; ok {
		return value
	}

	if start == len(numbers)-1 {
		return 1
	}

	result := 0

	end := len(numbers) - 1
	if start+3 < end {
		end = start + 3
	}
	for j := start + 1; j <= end; j++ {
		if numbers[j]-numbers[start] <= 3 {
			result += pathsToEnd(numbers, j)
		}
	}

	cache[start] = result
	return result
}
