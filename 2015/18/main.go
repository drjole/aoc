package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(first())
	fmt.Println(second())
}

func first() int {
	inputBytes, err := os.ReadFile("2015/18/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputBytes), "\n")
	grid := make([][]bool, len(lines))
	for i, line := range lines {
		grid[i] = make([]bool, len(line))
		for j, light := range line {
			switch light {
			case '#':
				grid[i][j] = true
			case '.':
				grid[i][j] = false
			default:
				panic("this should never happen")
			}
		}
	}

	const steps = 100
	for step := 0; step < steps; step++ {
		temp := make([][]bool, len(grid))
		for i := 0; i < len(grid); i++ {
			temp[i] = make([]bool, len(grid[i]))
			for j := 0; j < len(grid[i]); j++ {
				count := neighbours(grid, i, j)
				if grid[i][j] {
					if count == 2 || count == 3 {
						temp[i][j] = true
					} else {
						temp[i][j] = false
					}
				}
				if !grid[i][j] {
					if count == 3 {
						temp[i][j] = true
					} else {
						temp[i][j] = false
					}
				}
			}
		}
		grid = temp
	}

	on := 0
	for _, row := range grid {
		for _, light := range row {
			if light {
				on++
			}
		}
	}

	return on
}

func second() int {
	inputBytes, err := os.ReadFile("2015/18/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputBytes), "\n")
	grid := make([][]bool, len(lines))
	for i, line := range lines {
		grid[i] = make([]bool, len(line))
		for j, light := range line {
			switch light {
			case '#':
				grid[i][j] = true
			case '.':
				grid[i][j] = false
			default:
				panic("this should never happen")
			}
		}
	}

	const steps = 100
	for step := 0; step < steps; step++ {
		temp := make([][]bool, len(grid))

		for i := 0; i < len(grid); i++ {
			temp[i] = make([]bool, len(grid[i]))
		}

		temp[0][0] = true
		temp[0][len(grid[0])-1] = true
		temp[len(grid)-1][0] = true
		temp[len(grid)-1][len(grid[0])-1] = true

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if (i == 0 && j == 0) ||
					(i == 0 && j == len(grid[i])-1) ||
					(i == len(grid)-1 && j == 0) ||
					(i == len(grid)-1 && j == len(grid[i])-1) {
					continue
				}

				count := neighbours(grid, i, j)
				if grid[i][j] {
					if count == 2 || count == 3 {
						temp[i][j] = true
					} else {
						temp[i][j] = false
					}
				}
				if !grid[i][j] {
					if count == 3 {
						temp[i][j] = true
					} else {
						temp[i][j] = false
					}
				}
			}
		}

		grid = temp
	}

	on := 0
	for _, row := range grid {
		for _, light := range row {
			if light {
				on++
			}
		}
	}

	return on
}

func neighbours(grid [][]bool, i, j int) (result int) {
	for a := i - 1; a <= i+1; a++ {
		for b := j - 1; b <= j+1; b++ {
			center := a == i && b == j
			aInBounds := 0 <= a && a < len(grid)
			bInBounds := 0 <= b && b < len(grid[0])
			if !center && aInBounds && bInBounds {
				if grid[a][b] {
					result++
				}
			}
		}
	}
	return
}
