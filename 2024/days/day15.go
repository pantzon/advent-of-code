package days

import (
	"fmt"
	"path/filepath"
)

func d15Part1(path string) {}

func d15Part2(path string) {}

func Day15() {
	inputFile, err := filepath.Abs("./inputs/day15.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 15")
	d15Part1(inputFile)
	d15Part2(inputFile)
}
