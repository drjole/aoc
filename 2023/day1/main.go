package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func part1(input string) string {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		var first *int
		var last int
		for _, char := range line {
			var number *int
			if unicode.IsNumber(char) {
				n := int(char - '0')
				number = &n
			}
			if number != nil {
				if first == nil {
					first = number
				}
				last = *number
			}
		}
		n, _ := strconv.Atoi(fmt.Sprintf("%d%d", *first, last))
		sum += n
	}
	return fmt.Sprintf("%d", sum)
}

func part2(input string) string {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		var first *int
		var last int
		for i, char := range line {
			var number *int
			if unicode.IsNumber(char) {
				n := int(char - '0')
				number = &n
			} else if n, ok := spelledNumber(line[i:]); ok {
				number = &n
			}
			if number != nil {
				if first == nil {
					first = number
				}
				last = *number
			}
		}
		n, _ := strconv.Atoi(fmt.Sprintf("%d%d", *first, last))
		sum += n
	}
	return fmt.Sprintf("%d", sum)
}

func spelledNumber(s string) (int, bool) {
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	for word, number := range numbers {
		if strings.HasPrefix(s, word) {
			return number, true
		}
	}
	return 0, false
}

func main() {
	file, _ := os.ReadFile("input.txt")
	input := strings.TrimSuffix(string(file), "\n")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
