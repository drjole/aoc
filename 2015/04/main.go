package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("2015/04/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(first(input))
	fmt.Println(second(input))
}

func first(input []byte) int {
	i := 0
	for {
		hash := md5.Sum(append(input, []byte(fmt.Sprintf("%d", i))...))
		if fmt.Sprintf("%x", hash)[:5] == "00000" {
			return i
		}
		i++
	}
}

func second(input []byte) int {
	i := 0
	for {
		hash := md5.Sum(append(input, []byte(fmt.Sprintf("%d", i))...))
		if fmt.Sprintf("%x", hash)[:6] == "000000" {
			return i
		}
		i++
	}
}
