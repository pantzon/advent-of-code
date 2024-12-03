package days

import (
	"fmt"
	"path/filepath"
)

func d12Part1(path string) {}

func d12Part2(path string) {}

func Day12() {
	inputFile, err := filepath.Abs("./inputs/day12.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 12")
	d12Part1(inputFile)
	d12Part2(inputFile)
}
