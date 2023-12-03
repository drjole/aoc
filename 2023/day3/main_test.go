package main

import "testing"

func TestPart1(t *testing.T) {
	example := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	expected := "4361"
	if result := part1(example); result != expected {
		t.Fatalf("expected '%s' got '%s'", expected, result)
	}
}

func TestPart2(t *testing.T) {
	example := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	expected := "467835"
	if result := part2(example); result != expected {
		t.Fatalf("expected '%s' got '%s'", expected, result)
	}
}
