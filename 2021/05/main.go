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

type point struct {
	x, y int
}

func first() int {
	inputBytes, err := os.ReadFile("2021/05/input.txt")
	if err != nil {
		panic(err)
	}

	spots := make(map[point]int)
	for _, line := range strings.Split(string(inputBytes), "\n") {
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return r == ',' || r == '-' || r == ' '
		})
		x1, _ := strconv.Atoi(fields[0])
		y1, _ := strconv.Atoi(fields[1])
		x2, _ := strconv.Atoi(fields[3])
		y2, _ := strconv.Atoi(fields[4])
		switch {
		case x1 == x2:
			if y2 < y1 {
				y1, y2 = y2, y1
			}
			for i := y1; i <= y2; i++ {
				spots[point{x1, i}]++
			}
		case y1 == y2:
			if x2 < x1 {
				x1, x2 = x2, x1
			}
			for i := x1; i <= x2; i++ {
				spots[point{i, y1}]++
			}
		}
	}

	dangerousAreas := 0
	for _, p := range spots {
		if p >= 2 {
			dangerousAreas++
		}
	}

	return dangerousAreas
}

func second() int {
	inputBytes, err := os.ReadFile("2021/05/input.txt")
	if err != nil {
		panic(err)
	}

	spots := make(map[point]int)
	for _, line := range strings.Split(string(inputBytes), "\n") {
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return r == ',' || r == '-' || r == ' '
		})
		x1, _ := strconv.Atoi(fields[0])
		y1, _ := strconv.Atoi(fields[1])
		x2, _ := strconv.Atoi(fields[3])
		y2, _ := strconv.Atoi(fields[4])
		switch {
		case x1 == x2:
			if y2 < y1 {
				y1, y2 = y2, y1
			}
			for i := y1; i <= y2; i++ {
				spots[point{x1, i}]++
			}
		case y1 == y2:
			if x2 < x1 {
				x1, x2 = x2, x1
			}
			for i := x1; i <= x2; i++ {
				spots[point{i, y1}]++
			}
		default:
			xDir := 1
			if x2 < x1 {
				xDir = -1
			}
			yDir := 1
			if y2 < y1 {
				yDir = -1
			}
			for i := 0; float64(i) <= math.Abs(float64(x2-x1)); i++ {
				spots[point{x1 + i*xDir, y1 + i*yDir}]++
			}
		}
	}

	dangerousAreas := 0
	for _, p := range spots {
		if p >= 2 {
			dangerousAreas++
		}
	}

	return dangerousAreas
}
