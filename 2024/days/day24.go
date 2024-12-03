package days

import (
	"fmt"
	"path/filepath"
)

func d24Part1(path string) {}

func d24Part2(path string) {}

func Day24() {
	inputFile, err := filepath.Abs("./inputs/day24.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 24")
	d24Part1(inputFile)
	d24Part2(inputFile)
}
