package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(input string) string {
	sum := 0
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		buffer := line
		colon := strings.Index(buffer, ":")
		idString := buffer[strings.Index(buffer, " ")+1 : colon]
		id, _ := strconv.Atoi(idString)
		buffer = buffer[colon+2:]
		rounds := strings.Split(buffer, "; ")
		for _, round := range rounds {
			draws := strings.Split(round, ", ")
			for _, draw := range draws {
				components := strings.Split(draw, " ")
				countString := components[0]
				count, _ := strconv.Atoi(countString)
				color := components[1]
				if (color == "red" && count > 12) || (color == "green" && count > 13) || (color == "blue" && count > 14) {
					goto nextGame
				}
			}
		}
		sum += id
	nextGame:
	}
	return fmt.Sprintf("%d", sum)
}

func part2(input string) string {
	sum := 0
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		buffer := line
		colon := strings.Index(buffer, ":")
		buffer = buffer[colon+2:]
		rounds := strings.Split(buffer, "; ")
		minRed, minGreen, minBlue := 0, 0, 0
		for _, round := range rounds {
			draws := strings.Split(round, ", ")
			for _, draw := range draws {
				components := strings.Split(draw, " ")
				countString := components[0]
				count, _ := strconv.Atoi(countString)
				color := components[1]
				switch color {
				case "red":
					if count > minRed {
						minRed = count
					}
				case "green":
					if count > minGreen {
						minGreen = count
					}
				case "blue":
					if count > minBlue {
						minBlue = count
					}
				}
			}
		}
		sum += minRed * minGreen * minBlue
	}
	return fmt.Sprintf("%d", sum)
}

func main() {
	file, _ := os.ReadFile("input.txt")
	input := strings.TrimSuffix(string(file), "\n")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
