package days

import (
	"fmt"
	"path/filepath"
)

func d14Part1(path string) {}

func d14Part2(path string) {}

func Day14() {
	inputFile, err := filepath.Abs("./inputs/day14.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 14")
	d14Part1(inputFile)
	d14Part2(inputFile)
}
