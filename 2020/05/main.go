package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("2020/05/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(first(input))
	fmt.Println(second(input))
}

func first(input []byte) int {
	max := 0
	for _, inputString := range strings.Split(string(input), "\n") {
		rowBinary := strings.ReplaceAll(strings.ReplaceAll(inputString[:len(inputString)-3], "B", "1"), "F", "0")
		colBinary := strings.ReplaceAll(strings.ReplaceAll(inputString[len(inputString)-3:], "R", "1"), "L", "0")
		row, err := strconv.ParseInt(rowBinary, 2, 64)
		if err != nil {
			return -1
		}
		col, err := strconv.ParseInt(colBinary, 2, 64)
		if err != nil {
			return -1
		}
		seatId := int(row*8 + col)
		if seatId > max {
			max = seatId
		}
	}
	return max
}

func second(input []byte) int {
	mine := 0
	seats := make([]int, 0)
	for _, inputString := range strings.Split(string(input), "\n") {
		rowBinary := strings.ReplaceAll(strings.ReplaceAll(inputString[:len(inputString)-3], "B", "1"), "F", "0")
		colBinary := strings.ReplaceAll(strings.ReplaceAll(inputString[len(inputString)-3:], "R", "1"), "L", "0")
		row, err := strconv.ParseInt(rowBinary, 2, 64)
		if err != nil {
			return -1
		}
		if row == 0 || row == 127 {
			continue
		}
		col, err := strconv.ParseInt(colBinary, 2, 64)
		if err != nil {
			return -1
		}
		seats = append(seats, int(row*8+col))
	}
	sort.Slice(seats, func(i, j int) bool {
		return seats[i] < seats[j]
	})
	for i, seat := range seats {
		if i == 0 || i == len(seats)-1 {
			continue
		}
		if (seats[i-1] == seat-1) && (seats[i+1] == seat+1) {
			continue
		}
		return seat + 1
	}
	return mine
}
