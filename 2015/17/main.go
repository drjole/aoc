package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(first())
	fmt.Println(second())
}

func first() int {
	inputBytes, err := os.ReadFile("2015/17/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputBytes), "\n")
	l := len(lines)
	containers := make([]int, l)
	for i, line := range lines {
		value, _ := strconv.Atoi(line)
		containers[i] = value
	}

	const eggnog = 150
	num := 0
	for i := 1; i < (1 << l); i++ {
		capacity := 0
		for b := 0; b < l; b++ {
			if (i>>b)&1 == 1 {
				capacity += containers[b]
			}
		}
		if capacity == eggnog {
			num++
		}
	}

	return num
}

func second() int {
	inputBytes, err := os.ReadFile("2015/17/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputBytes), "\n")
	l := len(lines)
	containers := make([]int, l)
	for i, line := range lines {
		value, _ := strconv.Atoi(line)
		containers[i] = value
	}

	const eggnog = 150
	minimumNumberOfContainers := math.MaxInt
	containersUsed := make(map[int]int)
	for i := 1; i < (1 << l); i++ {
		capacity := 0
		numberOfContainers := 0
		for b := 0; b < l; b++ {
			if (i>>b)&1 == 1 {
				numberOfContainers++
				capacity += containers[b]
			}
		}
		if capacity == eggnog {
			containersUsed[numberOfContainers]++
			if numberOfContainers < minimumNumberOfContainers {
				minimumNumberOfContainers = numberOfContainers
			}
		}
	}

	return containersUsed[minimumNumberOfContainers]
}
