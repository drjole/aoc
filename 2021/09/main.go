package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(first())
	fmt.Println(second())
}

func first() int {
	file, err := os.Open("2021/09/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	numbers := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		j := 0
		numbers = append(numbers, make([]int, len(line)))
		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			numbers[i][j] = n
			j++
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	riskSum := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers[i]); j++ {
			noRisk := false
			for _, n := range neighbours(numbers, i, j) {
				if numbers[i][j] >= numbers[n.a][n.b] {
					noRisk = true
					break
				}
			}
			if noRisk {
				continue
			}
			riskSum += numbers[i][j] + 1
		}
	}

	return riskSum
}

func second() int {
	file, err := os.Open("2021/09/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	numbers := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		j := 0
		numbers = append(numbers, make([]int, len(line)))
		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			numbers[i][j] = n
			j++
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	basins := make([]int, 0)
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers[i]); j++ {
			noRisk := false
			for _, n := range neighbours(numbers, i, j) {
				if numbers[i][j] >= numbers[n.a][n.b] {
					noRisk = true
					break
				}
			}
			if noRisk {
				continue
			}

			// found a low point -> new basin
			v := make(map[point]struct{})
			visited := &v
			basinSize(numbers, visited, i, j)
			basins = append(basins, len(*visited))
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	return basins[0] * basins[1] * basins[2]
}

type point struct {
	x, y int
}

func basinSize(numbers [][]int, visited *map[point]struct{}, i, j int) {
	for _, n := range neighbours(numbers, i, j) {
		if numbers[n.a][n.b] == 9 {
			continue
		}
		if _, ok := (*visited)[point{n.a, n.b}]; ok {
			continue
		}
		(*visited)[point{n.a, n.b}] = struct{}{}
		basinSize(numbers, visited, n.a, n.b)
	}
	return
}

func neighbours(numbers [][]int, i, j int) (result []struct{ a, b int }) {
	for _, s := range []struct{ a, b int }{
		{a: i, b: j - 1},
		{a: i, b: j + 1},
		{a: i + 1, b: j},
		{a: i - 1, b: j},
	} {
		if s.a >= 0 && s.b >= 0 && s.a < len(numbers) && s.b < len(numbers[i]) {
			result = append(result, s)
		}
	}
	return
}
