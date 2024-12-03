package days

import (
	"fmt"
	"path/filepath"
)

func d21Part1(path string) {}

func d21Part2(path string) {}

func Day21() {
	inputFile, err := filepath.Abs("./inputs/day21.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 21")
	d21Part1(inputFile)
	d21Part2(inputFile)
}
