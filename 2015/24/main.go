package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile("2015/24/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputBytes), "\n")
	weights := make([]int, len(lines))
	for i, line := range lines {
		weights[i], _ = strconv.Atoi(line)
	}
	fmt.Println(minQuantumEntanglement(weights, sum(weights)/3))
	fmt.Println(minQuantumEntanglement(weights, sum(weights)/4))
}

func minQuantumEntanglement(weights []int, groupWeight int) int {
	combos := make([][]int, 0)
	for y := 1; y < len(weights); y++ {
		for _, c := range combinations(weights, y) {
			if sum(c) != groupWeight {
				continue
			}
			combos = append(combos, c)
		}
		if len(combos) > 0 {
			break
		}
	}

	qe := math.MaxInt
	for _, c := range combos {
		p := product(c)
		if p < qe {
			qe = p
		}
	}

	return qe
}

func combinations(set []int, groupSize int) [][]int {
	if groupSize == 1 {
		temp := make([][]int, 0)
		for _, element := range set {
			t := make([]int, 0)
			t = append(t, element)
			temp = append(temp, [][]int{t}...)
		}
		return temp
	}
	result := make([][]int, 0)
	for i := 0; i < len(set); i++ {
		permutations := make([]int, 0)
		permutations = append(permutations, set[:i]...)
		for _, x := range combinations(permutations, groupSize-1) {
			t := append(x, set[i])
			result = append(result, [][]int{t}...)
		}
	}
	return result
}

func sum(s []int) (r int) {
	for _, e := range s {
		r += e
	}
	return
}

func product(s []int) (r int) {
	r = 1
	for _, e := range s {
		r *= e
	}
	return
}
