package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println(first())
	fmt.Println(second())
}

func first() int {
	file, err := os.Open("2021/10/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := make(stack, 0)
		for _, char := range line {
			switch char {
			case '(', '[', '{', '<':
				s = s.push(char)
			case ')', ']', '}', '>':
				var n rune
				s, n = s.pop()
				if char != counterPart[n] {
					score += syntaxCheckerScore[char]
					break
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return score
}

func second() int {
	file, err := os.Open("2021/10/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	scores := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := make(stack, 0)
		score := 0
		line := scanner.Text()
		syntaxErrorFound := false
		for _, char := range line {
			switch char {
			case '(', '[', '{', '<':
				s = s.push(char)
			case ')', ']', '}', '>':
				var n rune
				s, n = s.pop()
				if char != counterPart[n] {
					syntaxErrorFound = true
					break
				}
			}
		}
		if !syntaxErrorFound {
			for len(s) > 0 {
				var c rune
				s, c = s.pop()
				score *= 5
				score += autocompleteScore[counterPart[c]]
			}
			scores = append(scores, score)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

var (
	counterPart        map[rune]rune
	syntaxCheckerScore map[rune]int
	autocompleteScore  map[rune]int
)

func init() {
	counterPart = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	syntaxCheckerScore = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	autocompleteScore = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
}

type stack []rune

func (s stack) push(n rune) stack {
	return append(s, n)
}

func (s stack) pop() (stack, rune) {
	if len(s) == 0 {
		return nil, 0
	}
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) string() string {
	b := strings.Builder{}
	for _, c := range s {
		b.WriteRune(c)
	}
	return b.String()
}
