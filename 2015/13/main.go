package main

import (
	"fmt"
	"github.com/drjole/aoc/util"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(first())
	fmt.Println(second())
}

func first() int {
	inputBytes, err := os.ReadFile("2015/13/input.txt")
	if err != nil {
		panic(err)
	}
	happiness := make(map[string]map[string]int)
	lines := strings.Split(string(inputBytes), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		value, _ := strconv.Atoi(fields[3])
		if fields[2] == "lose" {
			value *= -1
		}
		person := fields[0]
		if happiness[person] == nil {
			happiness[person] = make(map[string]int)
		}
		next := strings.ReplaceAll(fields[10], ".", "")
		happiness[person][next] = value
	}

	people := make([]string, len(happiness))
	i := 0
	for person := range happiness {
		people[i] = person
		i++
	}

	optimal := 0
	result := new([][]string)
	util.Permutations(result, people, len(people))
	for _, arrangement := range *result {
		total := 0

		for i, person := range arrangement {
			l := (i - 1) % len(people)
			if l < 0 {
				l = len(people) - 1
			}
			r := (i + 1) % len(people)
			if r > len(people)-1 {
				r = 0
			}
			left := arrangement[l]
			right := arrangement[r]
			total += happiness[person][left]
			total += happiness[person][right]
		}

		if total > optimal {
			optimal = total
		}
	}
	return optimal
}

func second() int {
	inputBytes, err := os.ReadFile("2015/13/input.txt")
	if err != nil {
		panic(err)
	}
	happiness := make(map[string]map[string]int)
	lines := strings.Split(string(inputBytes), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		value, _ := strconv.Atoi(fields[3])
		if fields[2] == "lose" {
			value *= -1
		}
		person := fields[0]
		if happiness[person] == nil {
			happiness[person] = make(map[string]int)
		}
		next := strings.ReplaceAll(fields[10], ".", "")
		happiness[person][next] = value
	}

	// Make sure that 'me' does not exist in the input list already.
	me := ""
	for {
		if _, ok := happiness[me]; !ok {
			break
		}
		me += "-"
	}

	happiness[me] = make(map[string]int)
	for person := range happiness {
		happiness[person][me] = 0
		happiness[me][person] = 0
	}

	people := make([]string, len(happiness))
	i := 0
	for person := range happiness {
		people[i] = person
		i++
	}

	optimal := 0
	result := new([][]string)
	util.Permutations(result, people, len(people))
	for _, arrangement := range *result {
		total := 0

		for i, person := range arrangement {
			l := (i - 1) % len(people)
			if l < 0 {
				l = len(people) - 1
			}
			r := (i + 1) % len(people)
			if r > len(people)-1 {
				r = 0
			}
			left := arrangement[l]
			right := arrangement[r]
			total += happiness[person][left]
			total += happiness[person][right]
		}

		if total > optimal {
			optimal = total
		}
	}
	return optimal
}
