package days

import (
	"fmt"
	"path/filepath"
)

func d8Part1(path string) {}

func d8Part2(path string) {}

func Day8() {
	inputFile, err := filepath.Abs("./inputs/day8.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 8")
	d8Part1(inputFile)
	d8Part2(inputFile)
}
