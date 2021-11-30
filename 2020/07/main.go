package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("2020/07/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(first(input))
	fmt.Println(second(input))
}

func first(input []byte) int {
	bags := make(map[string][]string)

	for _, inputString := range strings.Split(string(input), "\n") {
		splitted := strings.Split(inputString, " ")
		from := fmt.Sprintf("%s %s", splitted[0], splitted[1])
		if splitted[4] == "no" {
			bags[from] = []string{}
			continue
		}
		for len(splitted) > 4 {
			to := fmt.Sprintf("%s %s", splitted[5], splitted[6])
			bags[from] = append(bags[from], to)
			splitted = append(splitted[:4], splitted[8:]...)
		}
	}

	return numberOfBagsThatEventuallyContain("shiny gold", bags)
}

func numberOfBagsThatEventuallyContain(target string, bags map[string][]string) int {
	result := 0
	for color, innerBags := range bags {
		for _, innerBagColor := range innerBags {
			if innerBagColor == target {
				delete(bags, color)
				result += numberOfBagsThatEventuallyContain(color, bags) + 1
			}
		}
	}
	return result
}

type next struct {
	color string
	count int
}

func second(input []byte) int {
	bags := make(map[string][]next)

	for _, inputString := range strings.Split(string(input), "\n") {
		splitted := strings.Split(inputString, " ")
		from := fmt.Sprintf("%s %s", splitted[0], splitted[1])
		if splitted[4] == "no" {
			bags[from] = []next{}
			continue
		}
		for len(splitted) > 4 {
			to := fmt.Sprintf("%s %s", splitted[5], splitted[6])
			countString := splitted[4]
			count, err := strconv.Atoi(countString)
			if err != nil {
				panic(err)
			}
			bags[from] = append(bags[from], next{color: to, count: count})
			splitted = append(splitted[:4], splitted[8:]...)
		}
	}

	return numberOfBagsInside("shiny gold", bags)
}

func numberOfBagsInside(target string, bags map[string][]next) int {
	result := 0
	for _, innerBag := range bags[target] {
		result += innerBag.count * (numberOfBagsInside(innerBag.color, bags) + 1)
	}
	return result
}
