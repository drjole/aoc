package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(first())
	fmt.Println(second())
}

func first() int {
	inputBytes, err := os.ReadFile("2021/06/input.txt")
	if err != nil {
		panic(err)
	}

	initialPopulation := strings.Split(string(inputBytes), ",")
	population := make([]int, len(initialPopulation))
	for i := 0; i < len(initialPopulation); i++ {
		val, _ := strconv.Atoi(initialPopulation[i])
		population[i] = val
	}

	for day := 0; day < 80; day++ {
		populationCount := len(population)
		for i := 0; i < populationCount; i++ {
			if population[i] == 0 {
				population = append(population, 8)
				population[i] = 6
			} else {
				population[i]--
			}
		}
	}

	return len(population)
}

func second() int {
	inputBytes, err := os.ReadFile("2021/06/input.txt")
	if err != nil {
		panic(err)
	}

	initialPopulation := strings.Split(string(inputBytes), ",")
	population := make(map[int]int)
	for i := 0; i < len(initialPopulation); i++ {
		val, _ := strconv.Atoi(initialPopulation[i])
		population[val]++
	}

	for day := 0; day < 256; day++ {
		temp := population[0]
		for i := 0; i <= 7; i++ {
			population[i] = population[i+1]
		}
		population[6] += temp
		population[8] = temp
	}

	sum := 0
	for _, count := range population {
		sum += count
	}

	return sum
}
