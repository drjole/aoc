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
	file, err := os.Open("2015/19/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	molecules := make(map[string]struct{})
	var replacements [][2]string
	var molecule string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "=>") {
			molecule = line
			continue
		}
		fields := strings.Fields(line)
		replacements = append(replacements, [2]string{fields[0], fields[2]})
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for _, replacement := range replacements {
		for i := 0; i < len(molecule); i++ {
			if len(replacement[0]) > len(molecule)-i {
				break
			}
			if molecule[i:i+len(replacement[0])] == replacement[0] {
				newMolecule := molecule[:i] + replacement[1] + molecule[i+len(replacement[0]):]
				molecules[newMolecule] = struct{}{}
			}
		}
	}

	return len(molecules)
}

func second() int {
	file, err := os.Open("2015/19/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	replacements := make(map[string][]string)
	var molecule string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "=>") {
			molecule = line
			continue
		}
		fields := strings.Fields(line)
		if _, ok := replacements[fields[0]]; !ok {
			replacements[fields[0]] = make([]string, 0)
		}
		replacements[fields[2]] = append(replacements[fields[2]], fields[0])
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	steps := 0

	goal := molecule
	start := "e"

	for goal != start {
		for o, n := range replacements {
			if len(n) == 0 {
				continue
			}
			for {
				count := strings.Count(goal, o)
				if count <= 0 {
					break
				}
				steps += count
				goal = strings.ReplaceAll(goal, o, n[0])
			}
		}
	}

	return steps
}
