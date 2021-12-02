package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("2015/02/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(first(input))
	fmt.Println(second(input))
}

func first(input []byte) int {
	result := 0
	r := regexp.MustCompile(`(\d+)x(\d+)x(\d+)`)
	for _, inputString := range strings.Split(string(input), "\n") {
		match := r.FindStringSubmatch(inputString)
		lString, wString, hString := match[1], match[2], match[3]
		l, err := strconv.Atoi(lString)
		if err != nil {
			panic(err)
		}
		w, err := strconv.Atoi(wString)
		if err != nil {
			panic(err)
		}
		h, err := strconv.Atoi(hString)
		if err != nil {
			panic(err)
		}
		a, b, c := l*w, w*h, h*l
		min := a
		if b < min {
			min = b
		}
		if c < min {
			min = c
		}
		result += 2*a + 2*b + 2*c + min
	}
	return result
}

func second(input []byte) int {
	result := 0
	r := regexp.MustCompile(`(\d+)x(\d+)x(\d+)`)
	for _, inputString := range strings.Split(string(input), "\n") {
		match := r.FindStringSubmatch(inputString)
		lString, wString, hString := match[1], match[2], match[3]
		l, err := strconv.Atoi(lString)
		if err != nil {
			panic(err)
		}
		w, err := strconv.Atoi(wString)
		if err != nil {
			panic(err)
		}
		h, err := strconv.Atoi(hString)
		if err != nil {
			panic(err)
		}
		sides := []int{l, w, h}
		sort.Slice(sides, func(i, j int) bool {
			return sides[i] < sides[j]
		})
		result += 2*sides[0] + 2*sides[1] + sides[0]*sides[1]*sides[2]
	}
	return result
}
