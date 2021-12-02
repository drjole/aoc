package main

import (
	"bufio"
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
	file, err := os.Open("2021/02/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	horizontal := 0
	depth := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		direction, xString := line[0], line[1]
		x, _ := strconv.Atoi(xString)
		switch direction {
		case "forward":
			horizontal += x
		case "up":
			depth -= x
		case "down":
			depth += x
		default:
			panic("this should never happen")
		}
	}

	return horizontal * depth
}

func second() int {
	file, err := os.Open("2021/02/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	horizontal := 0
	depth := 0
	aim := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		direction, xString := line[0], line[1]
		x, _ := strconv.Atoi(xString)
		switch direction {
		case "forward":
			horizontal += x
			depth += aim * x
		case "up":
			aim -= x
		case "down":
			aim += x
		default:
			panic("this should never happen")
		}
	}

	return horizontal * depth
}
