package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("2020/08/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(first(input))
	fmt.Println(second(input))
}

func first(input []byte) int {
	visited := make(map[int]struct{})
	accumulator := 0
	programCounter := 0
	lines := strings.Split(string(input), "\n")

	for {
		if _, ok := visited[programCounter]; ok {
			break
		}
		visited[programCounter] = struct{}{}

		instruction := lines[programCounter]
		command := instruction[:3]
		valueString := instruction[4:]
		value, err := strconv.Atoi(valueString)
		if err != nil {
			panic(err)
		}

		switch command {
		case "acc":
			accumulator += value
			programCounter++
		case "jmp":
			programCounter += value
		case "nop":
			programCounter++
		default:
			panic("this should never happen")
		}
	}
	return accumulator
}

func second(input []byte) int {
	lines := strings.Split(string(input), "\n")
	for index, line := range lines {
		program := make([]string, len(lines))
		copy(program, lines)
		switch line[:3] {
		case "jmp":
			program[index] = strings.ReplaceAll(line, "jmp", "nop")
		case "nop":
			program[index] = strings.ReplaceAll(line, "nop", "jmp")
		default:
			continue
		}
		programCounter, accumulator := run(program)
		if programCounter > len(program)-1 {
			return accumulator
		}
	}

	return -1
}

func run(lines []string) (int, int) {
	visited := make(map[int]struct{})
	accumulator := 0
	programCounter := 0

	for {
		if programCounter > len(lines)-1 {
			break
		}
		if _, ok := visited[programCounter]; ok {
			break
		}
		visited[programCounter] = struct{}{}

		instruction := lines[programCounter]
		command := instruction[:3]
		valueString := instruction[4:]
		value, _ := strconv.Atoi(valueString)

		switch command {
		case "acc":
			accumulator += value
			programCounter++
		case "jmp":
			programCounter += value
		case "nop":
			programCounter++
		default:
			panic("this should never happen")
		}
	}
	return programCounter, accumulator
}
