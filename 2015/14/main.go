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
	inputBytes, err := os.ReadFile("2015/14/input.txt")
	if err != nil {
		panic(err)
	}
	type reindeer struct {
		speed     int
		endurance int
		rest      int
	}
	distances := make(map[reindeer]int)
	for _, line := range strings.Split(string(inputBytes), "\n") {
		fields := strings.Fields(line)
		speed, _ := strconv.Atoi(fields[3])
		endurance, _ := strconv.Atoi(fields[6])
		rest, _ := strconv.Atoi(fields[13])
		r := reindeer{
			speed:     speed,
			endurance: endurance,
			rest:      rest,
		}
		distances[r] = 0
	}

	for s := 0; s < 2503; s++ {
		for r := range distances {
			if s%(r.endurance+r.rest) < r.endurance {
				distances[r] += r.speed
			}
		}
	}

	winner := 0
	for _, distance := range distances {
		if distance > winner {
			winner = distance
		}
	}

	return winner
}

func second() int {
	inputBytes, err := os.ReadFile("2015/14/input.txt")
	if err != nil {
		panic(err)
	}
	type reindeer struct {
		speed     int
		endurance int
		rest      int
	}
	distances := make(map[reindeer]int)
	for _, line := range strings.Split(string(inputBytes), "\n") {
		fields := strings.Fields(line)
		speed, _ := strconv.Atoi(fields[3])
		endurance, _ := strconv.Atoi(fields[6])
		rest, _ := strconv.Atoi(fields[13])
		r := reindeer{
			speed:     speed,
			endurance: endurance,
			rest:      rest,
		}
		distances[r] = 0
	}

	points := make(map[reindeer]int)
	for s := 0; s < 2503; s++ {
		for r := range distances {
			if s%(r.endurance+r.rest) < r.endurance {
				distances[r] += r.speed
			}
		}

		leaders := 0
		for _, distance := range distances {
			if distance > leaders {
				leaders = distance
			}
		}

		for r, distance := range distances {
			if distance == leaders {
				points[r]++
			}
		}
	}

	best := 0
	for _, p := range points {
		if p > best {
			best = p
		}
	}

	return best
}
