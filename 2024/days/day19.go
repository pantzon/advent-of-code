package days

import (
	"fmt"
	"path/filepath"
)

func d19Part1(path string) {}

func d19Part2(path string) {}

func Day19() {
	inputFile, err := filepath.Abs("./inputs/day19.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 19")
	d19Part1(inputFile)
	d19Part2(inputFile)
}
