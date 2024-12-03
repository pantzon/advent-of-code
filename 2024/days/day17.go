package days

import (
	"fmt"
	"path/filepath"
)

func d17Part1(path string) {}

func d17Part2(path string) {}

func Day17() {
	inputFile, err := filepath.Abs("./inputs/day17.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 17")
	d17Part1(inputFile)
	d17Part2(inputFile)
}
