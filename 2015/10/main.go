package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(first())
	fmt.Println(second())
}

func first() int {
	numberBytes, err := os.ReadFile("2015/10/input.txt")
	if err != nil {
		panic(err)
	}
	number := string(numberBytes)
	for i := 0; i < 40; i++ {
		temp := ""
		count := 1
		previousChar := number[0]
		for _, char := range number[1:] {
			if uint8(char) == previousChar {
				count++
			} else if count > 0 {
				temp += fmt.Sprintf("%d%c", count, previousChar)
				previousChar = uint8(char)
				count = 1
			}
		}
		if count > 0 {
			temp += fmt.Sprintf("%d%c", count, previousChar)
		}
		number = temp
	}
	return len(number)
}

func second() int {
	numberBytes, err := os.ReadFile("2015/10/input.txt")
	if err != nil {
		panic(err)
	}
	number := string(numberBytes)
	for i := 0; i < 50; i++ {
		temp := strings.Builder{}
		count := 1
		previousChar := number[0]
		for _, char := range number[1:] {
			if uint8(char) == previousChar {
				count++
			} else if count > 0 {
				temp.WriteString(fmt.Sprintf("%d%c", count, previousChar))
				previousChar = uint8(char)
				count = 1
			}
		}
		if count > 0 {
			temp.WriteString(fmt.Sprintf("%d%c", count, previousChar))
		}
		number = temp.String()
	}
	return len(number)
}
