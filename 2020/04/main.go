package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("2020/04/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(first(input))
	fmt.Println(second(input))
}

func first(input []byte) int {
	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
	valid := 0
	for _, passport := range strings.Split(string(input), "\n\n") {
		checklist := make(map[string]struct{})
		rows := strings.Split(passport, "\n")
		checked := false
		for _, row := range rows {
			for _, entry := range strings.Split(row, " ") {
				for _, field := range requiredFields {
					if entry[:3] == field {
						checklist[field] = struct{}{}
						break
					}
				}
				if len(checklist) == len(requiredFields) {
					valid++
					checked = true
				}
				if checked {
					break
				}
			}
			if checked {
				break
			}
		}
	}
	return valid
}

func second(input []byte) int {
	hgtRegex := regexp.MustCompile(`^(\d+)(cm|in)$`)
	hclRegex := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	eclRegex := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	pidRegex := regexp.MustCompile(`^[0-9]{9}$`)

	validations := map[string]func(string) bool{
		"byr": func(s string) bool {
			if len(s) != 4 {
				return false
			}
			i, err := strconv.Atoi(s)
			if err != nil {
				return false
			}
			return 1920 <= i && i <= 2002
		},
		"iyr": func(s string) bool {
			if len(s) != 4 {
				return false
			}
			i, err := strconv.Atoi(s)
			if err != nil {
				return false
			}
			return 2010 <= i && i <= 2020
		},
		"eyr": func(s string) bool {
			if len(s) != 4 {
				return false
			}
			i, err := strconv.Atoi(s)
			if err != nil {
				return false
			}
			return 2020 <= i && i <= 2030
		},
		"hgt": func(s string) bool {
			match := hgtRegex.FindStringSubmatch(s)
			if len(match) != 3 {
				return false
			}
			heightString, unit := match[1], match[2]
			height, err := strconv.Atoi(heightString)
			if err != nil {
				return false
			}
			switch unit {
			case "cm":
				return 150 <= height && height <= 193
			case "in":
				return 59 <= height && height <= 76
			default:
				return false
			}
		},
		"hcl": func(s string) bool {
			return hclRegex.MatchString(s)
		},
		"ecl": func(s string) bool {
			return eclRegex.MatchString(s)
		},
		"pid": func(s string) bool {
			return pidRegex.MatchString(s)
		},
	}

	valid := 0
	for _, passport := range strings.Split(string(input), "\n\n") {
		checklist := make(map[string]struct{})
		rows := strings.Split(passport, "\n")
		checked := false
		for _, row := range rows {
			for _, entry := range strings.Split(row, " ") {
				for field, validation := range validations {
					if entry[:3] == field && validation(entry[4:]) {
						checklist[field] = struct{}{}
						break
					}
				}
				if len(checklist) == len(validations) {
					valid++
					checked = true
					break
				}
			}
			if checked {
				break
			}
		}
	}
	return valid
}
