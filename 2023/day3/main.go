package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type number struct {
	value   int
	indexes [][2]int
}

func part1(input string) string {
	lines := strings.Split(input, "\n")
	numbers := make([]number, 0)
	currentNumber := ""
	currentIndexes := make([][2]int, 0)
	for y, line := range lines {
		for x, char := range line {
			if unicode.IsDigit(char) {
				currentIndexes = append(currentIndexes, [2]int{x, y})
				currentNumber += string(char)
			} else if currentNumber != "" {
				value, _ := strconv.Atoi(currentNumber)
				numbers = append(numbers, number{value: value, indexes: currentIndexes})
				currentNumber = ""
				currentIndexes = make([][2]int, 0)
			}
		}
	}
	if currentNumber != "" {
		value, _ := strconv.Atoi(currentNumber)
		numbers = append(numbers, number{value: value, indexes: currentIndexes})
	}
	sum := 0
	for _, n := range numbers {
		for _, i := range n.indexes {
			neighs := neighborIndexes(input, i)
			for _, neigh := range neighs {
				char := rune(input[neigh[1]*(len(lines[0])+1)+neigh[0]])
				if !unicode.IsDigit(char) && char != '.' {
					sum += n.value
					goto nextNumber
				}
			}
		}
	nextNumber:
	}
	return fmt.Sprintf("%d", sum)
}

func part2(input string) string {
	lines := strings.Split(input, "\n")
	numbers := make([]number, 0)
	currentNumber := ""
	currentIndexes := make([][2]int, 0)
	for y, line := range lines {
		for x, char := range line {
			if unicode.IsDigit(char) {
				currentIndexes = append(currentIndexes, [2]int{x, y})
				currentNumber += string(char)
			} else if currentNumber != "" {
				value, _ := strconv.Atoi(currentNumber)
				numbers = append(numbers, number{value: value, indexes: currentIndexes})
				currentNumber = ""
				currentIndexes = make([][2]int, 0)
			}
		}
	}
	if currentNumber != "" {
		value, _ := strconv.Atoi(currentNumber)
		numbers = append(numbers, number{value: value, indexes: currentIndexes})
	}
	sum := 0
	for y, line := range lines {
		for x, char := range line {
			if char == '*' {
				neighs := neighborIndexes(input, [2]int{x, y})
				neighborCount := 0
				gearRatio := 1
				for _, n := range numbers {
					for _, i := range n.indexes {
						for _, neigh := range neighs {
							if neigh == i {
								neighborCount++
								gearRatio *= n.value
								goto nextNumber
							}
							if neighborCount > 2 {
								goto nextChar
							}
						}
					}
				nextNumber:
				}
				if neighborCount == 2 {
					sum += gearRatio
				}
			}
		nextChar:
		}
	}
	return fmt.Sprintf("%d", sum)
}

func neighborIndexes(engine string, index [2]int) [][2]int {
	neighbors := make([][2]int, 0)
	lines := strings.Split(engine, "\n")
	width := len(lines[0])
	height := len(lines)
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			newX := index[0] + x
			newY := index[1] + y
			if newX >= 0 && newX < width && newY >= 0 && newY < height {
				neighbors = append(neighbors, [2]int{newX, newY})
			}
		}
	}
	return neighbors
}

func main() {
	file, _ := os.ReadFile("input.txt")
	input := strings.TrimSuffix(string(file), "\n")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
