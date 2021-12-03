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

type coordinates struct {
	x, y int
}

func first() int {
	file, err := os.Open("2015/06/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	lights := make(map[coordinates]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		min := parseCoordinates(fields[len(fields)-3])
		max := parseCoordinates(fields[len(fields)-1])
		for x := min.x; x <= max.x; x++ {
			for y := min.y; y <= max.y; y++ {
				light := coordinates{x, y}
				if fields[0] == "toggle" {
					if _, ok := lights[light]; ok {
						delete(lights, light)
					} else {
						lights[light] = struct{}{}
					}
				} else if fields[1] == "on" {
					lights[light] = struct{}{}
				} else if fields[1] == "off" {
					delete(lights, light)
				} else {
					panic("this should never happen")
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return len(lights)
}

func second() int {
	file, err := os.Open("2015/06/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	lights := make(map[coordinates]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		min := parseCoordinates(fields[len(fields)-3])
		max := parseCoordinates(fields[len(fields)-1])
		for x := min.x; x <= max.x; x++ {
			for y := min.y; y <= max.y; y++ {
				light := coordinates{x, y}
				if fields[0] == "toggle" {
					lights[light] += 2
				} else if fields[1] == "on" {
					lights[light]++
				} else if fields[1] == "off" {
					if lights[light] > 0 {
						lights[light]--
					}
				} else {
					panic("this should never happen")
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	brightness := 0
	for _, light := range lights {
		brightness += light
	}
	return brightness
}

func parseCoordinates(c string) coordinates {
	splitted := strings.Split(c, ",")
	xString, yString := splitted[0], splitted[1]
	x, _ := strconv.Atoi(xString)
	y, _ := strconv.Atoi(yString)
	return coordinates{x, y}
}
