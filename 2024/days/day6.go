package days

import (
	"fmt"
	"path/filepath"
)

func d6Part1(path string) {}

func d6Part2(path string) {}

func Day6() {
	inputFile, err := filepath.Abs("./inputs/day6.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 6")
	d6Part1(inputFile)
	d6Part2(inputFile)
}
