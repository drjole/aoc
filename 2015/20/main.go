package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(first())
	fmt.Println(second())
}

func first() int {
	inputBytes, err := os.ReadFile("2015/20/input.txt")
	if err != nil {
		panic(err)
	}
	target, _ := strconv.Atoi(string(inputBytes))

	houseNumber := 1
	for {
		presents := 0
		for _, factor := range factors(houseNumber) {
			presents += 10 * factor
			if presents >= target {
				return houseNumber
			}
		}
		houseNumber++
	}
	return 0
}

func second() int {
	inputBytes, err := os.ReadFile("2015/20/input.txt")
	if err != nil {
		panic(err)
	}
	target, _ := strconv.Atoi(string(inputBytes))

	houseNumber := 1
	for {
		presents := 0
		for i, factor := range factors(houseNumber) {
			if i == 50 {
				break
			}
			presents += 11 * factor
			if presents >= target {
				return houseNumber
			}
		}
		houseNumber++
	}
	return 0
}

func factors(target int) (f []int) {
	i := 1
	for i*i < target {
		if target%i == 0 {
			if target/i == i {
				f = append(f, i)
			} else {
				f = append(f, i, target/i)
			}
		}
		i++
	}
	return
}
