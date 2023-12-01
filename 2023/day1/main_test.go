package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	example := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	expected := "142"
	if result := part1(example); result != expected {
		t.Fatalf("expected '%s' got '%s'", expected, result)
	}
}

func TestPart2(t *testing.T) {
	example := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	expected := "281"
	if result := part2(example); result != expected {
		t.Fatalf("expected '%s' go '%s'", expected, result)
	}
}
