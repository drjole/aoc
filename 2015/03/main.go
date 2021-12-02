package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("2015/03/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(first(input))
	fmt.Println(second(input))
}

func first(input []byte) int {
	var x, y int
	houses := make(map[string]struct{})
	for _, direction := range input {
		switch direction {
		case '>':
			x++
		case '<':
			x--
		case '^':
			y++
		case 'v':
			y--
		default:
			panic("this should never happen")
		}
		houses[fmt.Sprintf("%d,%d", x, y)] = struct{}{}
	}
	return len(houses)
}

func second(input []byte) int {
	var santaX, santaY, roboX, roboY *int
	santaX, santaY, roboX, roboY = new(int), new(int), new(int), new(int)
	*santaX, *santaY, *roboX, *roboY = 0, 0, 0, 0
	houses := make(map[string]struct{})
	for i, direction := range input {
		x, y := santaX, santaY
		if i%2 == 0 {
			x, y = roboX, roboY
		}
		switch direction {
		case '>':
			*x++
		case '<':
			*x--
		case '^':
			*y++
		case 'v':
			*y--
		default:
			panic("this should never happen")
		}
		houses[fmt.Sprintf("%d,%d", *x, *y)] = struct{}{}
	}
	return len(houses)
}
