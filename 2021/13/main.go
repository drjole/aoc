package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(first())
	second()
}

func first() int {
	inputBytes, err := os.ReadFile("2021/13/input.txt")
	if err != nil {
		panic(err)
	}

	splitted := strings.Split(string(inputBytes), "\n\n")
	dotStrings, foldStrings := splitted[0], splitted[1]

	dots := make(map[dot]struct{}, 0)
	for _, dotString := range strings.Split(dotStrings, "\n") {
		splitted := strings.Split(dotString, ",")
		xString, yString := splitted[0], splitted[1]
		x, _ := strconv.Atoi(xString)
		y, _ := strconv.Atoi(yString)
		dots[dot{x, y}] = struct{}{}
	}

	folds := make([]fold, 0)
	for _, foldString := range strings.Split(foldStrings, "\n") {
		axis := foldString[11]
		valueString := foldString[13:]
		value, _ := strconv.Atoi(valueString)
		folds = append(folds, fold{string(axis), value})
	}

	f := folds[0]
	for d := range dots {
		switch f.axis {
		case "x":
			if d.x > f.value {
				dots[dot{2*f.value - d.x, d.y}] = struct{}{}
				delete(dots, d)
			}
		case "y":
			if d.y > f.value {
				dots[dot{d.x, 2*f.value - d.y}] = struct{}{}
				delete(dots, d)
			}
		default:
			panic("this should never happen")
		}
	}

	return len(dots)
}

func second() {
	inputBytes, err := os.ReadFile("2021/13/input.txt")
	if err != nil {
		panic(err)
	}

	splitted := strings.Split(string(inputBytes), "\n\n")
	dotStrings, foldStrings := splitted[0], splitted[1]

	dots := make(map[dot]struct{}, 0)
	for _, dotString := range strings.Split(dotStrings, "\n") {
		splitted := strings.Split(dotString, ",")
		xString, yString := splitted[0], splitted[1]
		x, _ := strconv.Atoi(xString)
		y, _ := strconv.Atoi(yString)
		dots[dot{x, y}] = struct{}{}
	}

	folds := make([]fold, 0)
	for _, foldString := range strings.Split(foldStrings, "\n") {
		axis := foldString[11]
		valueString := foldString[13:]
		value, _ := strconv.Atoi(valueString)
		folds = append(folds, fold{string(axis), value})
	}

	for _, f := range folds {
		for d := range dots {
			switch f.axis {
			case "x":
				if d.x > f.value {
					dots[dot{2*f.value - d.x, d.y}] = struct{}{}
					delete(dots, d)
				}
			case "y":
				if d.y > f.value {
					dots[dot{d.x, 2*f.value - d.y}] = struct{}{}
					delete(dots, d)
				}
			default:
				panic("this should never happen")
			}
		}
	}

	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := 0, 0
	for d := range dots {
		if d.x < minX {
			minX = d.x
		}
		if d.y < minY {
			minY = d.y
		}
		if d.x > maxX {
			maxX = d.x
		}
		if d.y > maxY {
			maxY = d.y
		}
	}

	img := image.NewGray(image.Rectangle{Min: image.Point{X: minX, Y: minY}, Max: image.Point{X: maxX + 1, Y: maxY + 1}})
	for d := range dots {
		img.SetGray(d.x, d.y, color.Gray{Y: 255})
	}
	outFile, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := outFile.Close(); err != nil {
			panic(err)
		}
	}()
	if err := png.Encode(outFile, img); err != nil {
		panic(err)
	}
}

type dot struct {
	x, y int
}

type fold struct {
	axis  string
	value int
}
