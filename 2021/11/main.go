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
	inputBytes, err := os.ReadFile("2021/11/input.txt")
	if err != nil {
		panic(err)
	}

	octopuses := make([][]int, 0)
	i := 0
	for _, line := range strings.Split(string(inputBytes), "\n") {
		octopuses = append(octopuses, make([]int, len(line)))
		j := 0
		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			octopuses[i][j] = n
			j++
		}
		i++
	}

	totalFlashes := 0
	for step := 0; step < 100; step++ {

		// Keep track of the octopuses that already flashed in this step.
		flashes := make(map[struct{ x, y int }]struct{})

		// First, the energy level of each octopus increases by 1.
		for i := 0; i < len(octopuses); i++ {
			for j := 0; j < len(octopuses[i]); j++ {
				octopuses[i][j]++
			}
		}

	startOver:
		for i := 0; i < len(octopuses); i++ {
			for j := 0; j < len(octopuses[i]); j++ {
				if flash(&octopuses, i, j, &flashes) {
					goto startOver
				}
			}
		}

		totalFlashes += len(flashes)

		for f := range flashes {
			octopuses[f.x][f.y] = 0
		}
	}

	return totalFlashes
}

func second() int {
	inputBytes, err := os.ReadFile("2021/11/input.txt")
	if err != nil {
		panic(err)
	}

	octopuses := make([][]int, 0)
	i := 0
	for _, line := range strings.Split(string(inputBytes), "\n") {
		octopuses = append(octopuses, make([]int, len(line)))
		j := 0
		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			octopuses[i][j] = n
			j++
		}
		i++
	}

	step := 1
	for {

		// Keep track of the octopuses that already flashed in this step.
		flashes := make(map[struct{ x, y int }]struct{})

		// First, the energy level of each octopus increases by 1.
		for i := 0; i < len(octopuses); i++ {
			for j := 0; j < len(octopuses[i]); j++ {
				octopuses[i][j]++
			}
		}

	startOver:
		for i := 0; i < len(octopuses); i++ {
			for j := 0; j < len(octopuses[i]); j++ {
				if flash(&octopuses, i, j, &flashes) {
					goto startOver
				}
			}
		}

		for f := range flashes {
			octopuses[f.x][f.y] = 0
		}

		if len(flashes) == len(octopuses)*len(octopuses[0]) {
			return step
		}
		step++
	}

	return 0
}

func flash(octopuses *[][]int, i, j int, flashes *map[struct{ x, y int }]struct{}) bool {
	if _, ok := (*flashes)[struct{ x, y int }{i, j}]; ok {
		return false
	}

	if (*octopuses)[i][j] <= 9 {
		return false
	}

	(*flashes)[struct{ x, y int }{i, j}] = struct{}{}

	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x == i && y == j {
				continue
			}
			if x < 0 || y < 0 || x > len(*octopuses)-1 || y > len((*octopuses)[i])-1 {
				continue
			}
			(*octopuses)[x][y]++
		}
	}

	return true
}
