package days

import (
	"fmt"
	"path/filepath"
)

func d18Part1(path string) {}

func d18Part2(path string) {}

func Day18() {
	inputFile, err := filepath.Abs("./inputs/day18.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 18")
	d18Part1(inputFile)
	d18Part2(inputFile)
}
