package days

import (
	"fmt"
	"path/filepath"
)

func d20Part1(path string) {}

func d20Part2(path string) {}

func Day20() {
	inputFile, err := filepath.Abs("./inputs/day20.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 20")
	d20Part1(inputFile)
	d20Part2(inputFile)
}
