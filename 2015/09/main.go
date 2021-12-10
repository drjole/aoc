package main

import (
	"bufio"
	"fmt"
	"github.com/drjole/aoc/util"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(first())
	fmt.Println(second())
}

type edge struct {
	from     string
	to       string
	distance int
}

func first() int {
	file, err := os.Open("2015/09/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(file)

	graph := make(map[string][]*edge)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		from, to, distanceString := fields[0], fields[2], fields[4]
		distance, _ := strconv.Atoi(distanceString)
		graph[from] = append(graph[from], &edge{from, to, distance})
		graph[to] = append(graph[to], &edge{to, from, distance})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var node string
	for n := range graph {
		if node == "" {
			node = n
			break
		}
	}
	keys := make([]string, 0)
	for key := range graph {
		keys = append(keys, key)
	}
	smallestDistance := math.MaxInt
	result := new([][]string)
	util.Permutations(result, keys, len(keys))
	for _, permutation := range *result {
		distance := 0
		for i, city := range permutation[:len(permutation)-1] {
			for _, edge := range graph[city] {
				if edge.to == permutation[i+1] {
					distance += edge.distance
					break
				}
			}
		}
		if distance < smallestDistance {
			smallestDistance = distance
		}
	}
	return smallestDistance
}

func second() int {
	file, err := os.Open("2015/09/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(file)

	graph := make(map[string][]*edge)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		from, to, distanceString := fields[0], fields[2], fields[4]
		distance, _ := strconv.Atoi(distanceString)
		graph[from] = append(graph[from], &edge{from, to, distance})
		graph[to] = append(graph[to], &edge{to, from, distance})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var node string
	for n := range graph {
		if node == "" {
			node = n
			break
		}
	}
	keys := make([]string, 0)
	for key := range graph {
		keys = append(keys, key)
	}
	longestDistance := 0
	result := new([][]string)
	util.Permutations(result, keys, len(keys))
	for _, permutation := range *result {
		distance := 0
		for i, city := range permutation[:len(permutation)-1] {
			for _, edge := range graph[city] {
				if edge.to == permutation[i+1] {
					distance += edge.distance
					break
				}
			}
		}
		if distance > longestDistance {
			longestDistance = distance
		}
	}
	return longestDistance
}
