package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(first())
	fmt.Println(second())
}

func first() int {
	file, err := os.Open("2015/08/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(file)

	codeCharacters := 0
	inMemoryCharacters := 0
	for scanner.Scan() {
		line := scanner.Text()
		codeCharacters += len(line)
		line = line[1 : len(line)-1]
		for i := 0; i < len(line); i++ {
			if line[i] != '\\' {
				continue
			}
			switch line[i+1] {
			case '\\', '"':
				line = line[:i] + line[i+1:]
			case 'x':
				line = line[:i] + "0" + line[i+4:]
			}
		}
		inMemoryCharacters += len(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return codeCharacters - inMemoryCharacters
}

func second() int {
	file, err := os.Open("2015/08/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(file)

	codeCharacters := 0
	escapedCodeCharacters := 0
	for scanner.Scan() {
		line := scanner.Text()
		codeCharacters += len(line)
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '\\', '"':
				line = line[:i] + "\\" + line[i:]
				i++
			}
		}
		escapedCodeCharacters += len(line) + 2
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return escapedCodeCharacters - codeCharacters
}
