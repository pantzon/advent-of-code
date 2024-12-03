package days

import (
	"fmt"
	"path/filepath"
)

func d11Part1(path string) {}

func d11Part2(path string) {}

func Day11() {
	inputFile, err := filepath.Abs("./inputs/day11.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 11")
	d11Part1(inputFile)
	d11Part2(inputFile)
}
