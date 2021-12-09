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

var cache map[string]uint16

func init() {
	cache = make(map[string]uint16)
}

func first() uint16 {
	inputBytes, err := os.ReadFile("2015/07/input.txt")
	if err != nil {
		panic(err)
	}
	gates := make(map[string][]string)

	lines := strings.Split(string(inputBytes), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		l := len(fields)
		gates[fields[l-1]] = fields[:l-2]
	}

	return evaluate(gates, "a")
}

func second() uint16 {
	inputBytes, err := os.ReadFile("2015/07/input.txt")
	if err != nil {
		panic(err)
	}
	gates := make(map[string][]string)

	lines := strings.Split(string(inputBytes), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		l := len(fields)
		gates[fields[l-1]] = fields[:l-2]
	}

	a := evaluate(gates, "a")
	cache = make(map[string]uint16)
	cache["b"] = a
	return evaluate(gates, "a")
}

func evaluate(gates map[string][]string, wire string) uint16 {
	if cached, ok := cache[wire]; ok {
		return cached
	}

	value, err := strconv.Atoi(wire)
	if err == nil {
		return uint16(value)
	}

	gate := gates[wire]

	switch len(gate) {
	case 1:
		return evaluate(gates, gate[0])
	case 2:
		return ^evaluate(gates, gate[1])
	case 3:
		switch gate[1] {
		case "AND":
			cache[wire] = evaluate(gates, gate[0]) & evaluate(gates, gate[2])
		case "OR":
			cache[wire] = evaluate(gates, gate[0]) | evaluate(gates, gate[2])
		case "LSHIFT":
			shift, _ := strconv.Atoi(gate[2])
			cache[wire] = evaluate(gates, gate[0]) << shift
		case "RSHIFT":
			shift, _ := strconv.Atoi(gate[2])
			cache[wire] = evaluate(gates, gate[0]) >> shift
		}
	default:
		panic("this should never happen")
	}
	return cache[wire]
}
