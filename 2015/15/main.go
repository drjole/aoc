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

func first() int {
	inputBytes, err := os.ReadFile("2015/15/input.txt")
	if err != nil {
		panic(err)
	}

	ingredients := make([][5]int, 0)
	for _, line := range strings.Split(string(inputBytes), "\n") {
		fields := strings.Fields(line)
		ingredient := [5]int{}
		for i := 0; i < 5; i++ {
			val, _ := strconv.Atoi(strings.ReplaceAll(fields[2*i+2], ",", ""))
			ingredient[i] = val
		}
		ingredients = append(ingredients, ingredient)
	}
	best := 0
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			for c := 0; c < 100; c++ {
				for d := 0; d < 100; d++ {
					if a+b+c+d != 100 {
						continue
					}
					score := 1
					for i := 0; i < 4; i++ {
						s := a*ingredients[0][i] + b*ingredients[1][i] + c*ingredients[2][i] + d*ingredients[3][i]
						if s < 0 {
							score = 0
							break
						}
						score *= s
					}
					if score > best {
						best = score
					}
				}
			}
		}
	}

	return best
}

func second() int {
	inputBytes, err := os.ReadFile("2015/15/input.txt")
	if err != nil {
		panic(err)
	}

	ingredients := make([][5]int, 0)
	for _, line := range strings.Split(string(inputBytes), "\n") {
		fields := strings.Fields(line)
		ingredient := [5]int{}
		for i := 0; i < 5; i++ {
			val, _ := strconv.Atoi(strings.ReplaceAll(fields[2*i+2], ",", ""))
			ingredient[i] = val
		}
		ingredients = append(ingredients, ingredient)
	}
	best := 0
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			for c := 0; c < 100; c++ {
				for d := 0; d < 100; d++ {
					if a+b+c+d != 100 {
						continue
					}
					calories := a*ingredients[0][4] + b*ingredients[1][4] + c*ingredients[2][4] + d*ingredients[3][4]
					if calories != 500 {
						continue
					}
					score := 1
					for i := 0; i < 4; i++ {
						s := a*ingredients[0][i] + b*ingredients[1][i] + c*ingredients[2][i] + d*ingredients[3][i]
						if s < 0 {
							score = 0
							break
						}
						score *= s
					}
					if score > best {
						best = score
					}
				}
			}
		}
	}

	return best
}
