package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile("2015/25/input.txt")
	if err != nil {
		panic(err)
	}
	fields := strings.Fields(string(inputBytes))
	targetRow, _ := strconv.Atoi(strings.ReplaceAll(fields[15], ",", ""))
	targetCol, _ := strconv.Atoi(strings.ReplaceAll(fields[17], ".", ""))

	row := 1
	col := 1
	var prevN int
	n := 20151125
	for {
		row = col
		col = 1
		for row >= 1 {
			if row == targetRow && col == targetCol {
				fmt.Println(n)
				return
			}
			prevN = n
			n = (prevN * 252533) % 33554393
			row--
			col++
		}
	}
}
