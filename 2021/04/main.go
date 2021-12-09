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

// N is the number of rows/cols in a grid
const N = 5

func first() int {
	inputBytes, err := os.ReadFile("2021/04/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputBytes), "\n\n")
	numberStrings := strings.Split(lines[0], ",")
	gridStrings := lines[1:]

	grids := make([][N * N]int, len(gridStrings))

	for i, gridString := range gridStrings {
		j := 0
		for _, gridStringLine := range strings.Split(gridString, "\n") {
			for _, gridEntry := range strings.Split(gridStringLine, " ") {
				if gridEntry == "" {
					continue
				}
				value, _ := strconv.Atoi(gridEntry)
				grids[i][j] = value
				j++
			}
		}
	}

	for _, numberString := range numberStrings {
		number, _ := strconv.Atoi(numberString)
		for gridIndex := 0; gridIndex < len(grids); gridIndex++ {
			for gridEntryIndex := 0; gridEntryIndex < N*N; gridEntryIndex++ {
				if grids[gridIndex][gridEntryIndex] != number {
					continue
				}

				// We got a bingo
				grids[gridIndex][gridEntryIndex] = -1

				if won(grids[gridIndex]) {
					sum := 0
					for i := 0; i < N*N; i++ {
						if grids[gridIndex][i] != -1 {
							sum += grids[gridIndex][i]
						}
					}
					return number * sum
				}
			}
		}
	}

	return 0
}

func second() int {
	inputBytes, err := os.ReadFile("2021/04/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputBytes), "\n\n")
	numberStrings := strings.Split(lines[0], ",")
	gridStrings := lines[1:]

	grids := make([][N * N]int, len(gridStrings))

	for i, gridString := range gridStrings {
		j := 0
		for _, gridStringLine := range strings.Split(gridString, "\n") {
			for _, gridEntry := range strings.Split(gridStringLine, " ") {
				if gridEntry == "" {
					continue
				}
				value, _ := strconv.Atoi(gridEntry)
				grids[i][j] = value
				j++
			}
		}
	}

	winners := make(map[int]struct{})
	lastWinnerNumber := -1
	lastWinner := -1
	for _, numberString := range numberStrings {
		number, _ := strconv.Atoi(numberString)
		for gridIndex := 0; gridIndex < len(grids); gridIndex++ {
			if _, ok := winners[gridIndex]; ok {
				continue
			}
			for gridEntryIndex := 0; gridEntryIndex < N*N; gridEntryIndex++ {
				if grids[gridIndex][gridEntryIndex] != number {
					continue
				}

				// We got a bingo
				grids[gridIndex][gridEntryIndex] = -1

				if won(grids[gridIndex]) {
					winners[gridIndex] = struct{}{}
					lastWinnerNumber = number
					lastWinner = gridIndex
				}
			}
		}
	}

	sum := 0
	for i := 0; i < N*N; i++ {
		if grids[lastWinner][i] != -1 {
			sum += grids[lastWinner][i]
		}
	}
	return lastWinnerNumber * sum
}

func won(grid [N * N]int) bool {
	for i := 0; i < N; i++ {
		rowWins := true
		columnWins := true
		for j := 0; j < N; j++ {
			if grid[N*i+j] != -1 {
				rowWins = false
			}
			if grid[N*j+i] != -1 {
				columnWins = false
			}
			if !rowWins && !columnWins {
				break
			}
		}
		if rowWins || columnWins {
			return true
		}
	}
	return false
}
