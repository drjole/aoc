package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(first())
	fmt.Println(second())
}

func first() int {
	file, err := os.Open("2015/05/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	nice := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		invalid := false
		line := scanner.Text()
		for _, forbidden := range []string{"ab", "cd", "pq", "xy"} {
			if strings.Contains(line, forbidden) {
				invalid = true
				break
			}
		}
		if invalid {
			continue
		}

		vowels := 0
		doubles := false
		var previous uint8
		for _, char := range line {
			switch uint8(char) {
			case 'a', 'e', 'i', 'o', 'u':
				vowels++
			}
			if uint8(char) == previous {
				doubles = true
			}
			if vowels >= 3 && doubles {
				nice++
				break
			}
			previous = uint8(char)
		}

	}
	return nice
}

func second() int {
	file, err := os.Open("2015/05/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	nice := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		enclosing := false
		doubleDoubles := false
		for i := 0; i < len(line); i++ {
			if i < len(line)-2 && line[i] == line[i+2] {
				enclosing = true
			}
			if i < len(line)-3 && strings.Contains(line[i+2:], line[i:i+2]) {
				doubleDoubles = true
			}
			if enclosing && doubleDoubles {
				nice++
				break
			}
		}
	}
	return nice
}
