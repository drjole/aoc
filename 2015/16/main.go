package main

import (
	"bufio"
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
	var samples = map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	file, err := os.Open("2015/16/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	sue := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		found := true
		line := scanner.Text()
		fields := strings.Fields(line)
		for i := 2; i < len(fields); i += 2 {
			thing := strings.ReplaceAll(fields[i], ":", "")
			countString := strings.ReplaceAll(fields[i+1], ",", "")
			count, _ := strconv.Atoi(countString)
			if count != samples[thing] {
				found = false
				sue++
				break
			}
		}
		if found {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return sue
}

func second() int {
	var samples = map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	file, err := os.Open("2015/16/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	sue := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		found := true
		line := scanner.Text()
		fields := strings.Fields(line)
		for i := 2; i < len(fields); i += 2 {
			thing := strings.ReplaceAll(fields[i], ":", "")
			countString := strings.ReplaceAll(fields[i+1], ",", "")
			count, _ := strconv.Atoi(countString)
			switch thing {
			case "cats", "trees":
				if count <= samples[thing] {
					found = false
					sue++
					break
				}
			case "pomeranians", "goldfish":
				if count >= samples[thing] {
					found = false
					sue++
					break
				}
			default:
				if count != samples[thing] {
					found = false
					sue++
					break
				}
			}
			if !found {
				break
			}
		}
		if found {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return sue
}
