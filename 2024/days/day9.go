package days

import (
	"fmt"
	"path/filepath"
)

func d9Part1(path string) {}

func d9Part2(path string) {}

func Day9() {
	inputFile, err := filepath.Abs("./inputs/day9.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 9")
	d9Part1(inputFile)
	d9Part2(inputFile)
}
