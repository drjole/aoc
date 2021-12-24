package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(first())
	fmt.Println(second())
}

func first() int {
	inputBytes, err := os.ReadFile("2021/07/input.txt")
	if err != nil {
		panic(err)
	}

	crabs := make([]int, 0)
	for _, s := range strings.Split(string(inputBytes), ",") {
		val, _ := strconv.Atoi(s)
		crabs = append(crabs, val)
	}

	sort.Ints(crabs)

	minFuel := float64(math.MaxInt)
	for target := crabs[0]; target <= crabs[len(crabs)-1]; target++ {
		fuel := 0.0
		for _, c := range crabs {
			fuel += math.Abs(float64(c - target))
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	return int(minFuel)
}

func second() int {
	inputBytes, err := os.ReadFile("2021/07/input.txt")
	if err != nil {
		panic(err)
	}

	crabs := make([]int, 0)
	for _, s := range strings.Split(string(inputBytes), ",") {
		val, _ := strconv.Atoi(s)
		crabs = append(crabs, val)
	}

	sort.Ints(crabs)

	minFuel := math.MaxInt
	for target := crabs[0]; target <= crabs[len(crabs)-1]; target++ {
		fuel := 0
		for _, c := range crabs {
			f := math.Abs(float64(c - target))
			fuel += int((f*f + f) / 2)
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	return minFuel
}
