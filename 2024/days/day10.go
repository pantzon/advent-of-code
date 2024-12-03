package days

import (
	"fmt"
	"path/filepath"
)

func d10Part1(path string) {}

func d10Part2(path string) {}

func Day10() {
	inputFile, err := filepath.Abs("./inputs/day10.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 10")
	d10Part1(inputFile)
	d10Part2(inputFile)
}
