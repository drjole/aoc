package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println(first())
	fmt.Println(second())
}

func first() int {
	file, err := os.Open("2015/12/input.txt")
	if err != nil {
		panic(err)
	}

	sum := float64(0)
	decoder := json.NewDecoder(file)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if number, ok := t.(float64); ok {
			sum += number
		}
	}
	return int(sum)
}

func second() int {
	inputBytes, err := os.ReadFile("2015/12/input.txt")
	if err != nil {
		panic(err)
	}

	var buffer interface{}
	if err := json.Unmarshal(inputBytes, &buffer); err != nil {
		panic(err)
	}
	return int(sum(buffer))
}

func sum(b interface{}) float64 {
	total := float64(0)
	switch v := b.(type) {
	case float64:
		total = v
	case []interface{}:
		for _, value := range v {
			total += sum(value)
		}
	case map[string]interface{}:
		for _, value := range v {
			if value == "red" {
				total = 0
				break
			}
			total += sum(value)
		}
	}
	return total
}
