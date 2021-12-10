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

func first() uint {
	inputBytes, err := os.ReadFile("2015/23/input.txt")
	if err != nil {
		panic(err)
	}
	instructionStrings := strings.Split(string(inputBytes), "\n")

	a, b, pc := new(uint), new(uint), new(uint)
	instructions := make([]instruction, 0)
	maxPc := uint(len(instructionStrings) - 1)

	for _, instructionString := range instructionStrings {
		fields := strings.Fields(instructionString)
		i := instruction{
			name: fields[0],
		}

		switch fields[0] {
		case "hlf", "tpl", "inc", "jie", "jio":
			if strings.ReplaceAll(fields[1], ",", "") == "a" {
				i.r = a
			} else {
				i.r = b
			}
		}

		switch fields[0] {
		case "jio", "jie":
			offsetString := fields[2]
			offset, _ := strconv.Atoi(offsetString)
			i.offset = offset
		case "jmp":
			offsetString := fields[1]
			offset, _ := strconv.Atoi(offsetString)
			i.offset = offset
		}

		instructions = append(instructions, i)
	}

	for {
		if *pc > maxPc {
			break
		}
		i := instructions[*pc]
		switch i.name {
		case "hlf":
			hlf(pc, i.r)
		case "tpl":
			tpl(pc, i.r)
		case "inc":
			inc(pc, i.r)
		case "jmp":
			jmp(pc, i.offset)
		case "jie":
			jie(pc, i.r, i.offset)
		case "jio":
			jio(pc, i.r, i.offset)
		}
	}

	return *b
}

func second() uint {
	inputBytes, err := os.ReadFile("2015/23/input.txt")
	if err != nil {
		panic(err)
	}
	instructionStrings := strings.Split(string(inputBytes), "\n")

	a, b, pc := new(uint), new(uint), new(uint)
	*a = 1
	instructions := make([]instruction, 0)
	maxPc := uint(len(instructionStrings) - 1)

	for _, instructionString := range instructionStrings {
		fields := strings.Fields(instructionString)
		i := instruction{
			name: fields[0],
		}

		switch fields[0] {
		case "hlf", "tpl", "inc", "jie", "jio":
			if strings.ReplaceAll(fields[1], ",", "") == "a" {
				i.r = a
			} else {
				i.r = b
			}
		}

		switch fields[0] {
		case "jio", "jie":
			offsetString := fields[2]
			offset, _ := strconv.Atoi(offsetString)
			i.offset = offset
		case "jmp":
			offsetString := fields[1]
			offset, _ := strconv.Atoi(offsetString)
			i.offset = offset
		}

		instructions = append(instructions, i)
	}

	for {
		if *pc > maxPc {
			break
		}
		i := instructions[*pc]
		switch i.name {
		case "hlf":
			hlf(pc, i.r)
		case "tpl":
			tpl(pc, i.r)
		case "inc":
			inc(pc, i.r)
		case "jmp":
			jmp(pc, i.offset)
		case "jie":
			jie(pc, i.r, i.offset)
		case "jio":
			jio(pc, i.r, i.offset)
		}
	}

	return *b
}

type instruction struct {
	name   string
	r      *uint
	offset int
}

func hlf(pc, r *uint) {
	*r /= 2
	*pc++
}

func tpl(pc, r *uint) {
	*r *= 3
	*pc++
}

func inc(pc, r *uint) {
	*r++
	*pc++
}

func jmp(pc *uint, offset int) {
	*pc += uint(offset)
}

func jie(pc, r *uint, offset int) {
	if *r%2 == 0 {
		jmp(pc, offset)
	} else {
		*pc++
	}
}

func jio(pc, r *uint, offset int) {
	if *r == 1 {
		jmp(pc, offset)
	} else {
		*pc++
	}
}
