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

func first() int64 {
	inputBytes, err := os.ReadFile("2021/03/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputBytes), "\n")

	var bits []int
	for _, line := range lines {
		if len(bits) == 0 {
			bits = make([]int, len(line))
		}
		for i, bit := range line {
			if bit == '1' {
				bits[i]++
			} else {
				bits[i]--
			}
		}
	}

	gammaBits := ""
	epsilonBits := ""
	for _, bit := range bits {
		if bit > 0 {
			gammaBits += "1"
			epsilonBits += "0"
		} else {
			gammaBits += "0"
			epsilonBits += "1"
		}
	}

	gamma, err := strconv.ParseInt(gammaBits, 2, 64)
	if err != nil {
		panic(err)
	}
	epsilon, err := strconv.ParseInt(epsilonBits, 2, 64)
	if err != nil {
		panic(err)
	}

	return gamma * epsilon
}

func second() int64 {
	inputBytes, err := os.ReadFile("2021/03/input.txt")
	if err != nil {
		panic(err)
	}

	oxygenLines := strings.Split(string(inputBytes), "\n")
	carbonLines := strings.Split(string(inputBytes), "\n")
	oxygenBitPos := 0
	carbonBitPos := 0
	for {
		if len(oxygenLines) > 1 {
			bit := uint8(0)
			if countBits(oxygenLines, oxygenBitPos) >= 0 {
				bit = 1
			}
			filter(&oxygenLines, oxygenBitPos, bit)
			oxygenBitPos++
		}
		if len(carbonLines) > 1 {
			bit := uint8(0)
			if countBits(carbonLines, carbonBitPos) < 0 {
				bit = 1
			}
			filter(&carbonLines, carbonBitPos, bit)
			carbonBitPos++
		}
		if len(oxygenLines) == 1 && len(carbonLines) == 1 {
			break
		}
	}

	oxygen, _ := strconv.ParseInt(oxygenLines[0], 2, 64)
	carbon, _ := strconv.ParseInt(carbonLines[0], 2, 64)

	return oxygen * carbon
}

func countBits(lines []string, bitPos int) int {
	counter := 0
	for _, line := range lines {
		if line[bitPos] == '1' {
			counter++
		} else {
			counter--
		}
	}
	return counter
}

func filter(lines *[]string, bitPos int, bit uint8) {
	for i := 0; i < len(*lines); i++ {
		if (*lines)[i][bitPos]-48 != bit {
			*lines = append((*lines)[:i], (*lines)[i+1:]...)
			i -= 1
			if len(*lines) == 1 {
				return
			}
		}
	}
}
